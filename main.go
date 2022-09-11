package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/memphisdev/memphis.go"
)

func getMsgInSize(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(1)
	}
	return bytes
}

func main() {
	opType := (strings.Split(os.Args[1], "="))[1]
	if opType != "produce" && opType != "consume" {
		fmt.Println("opType has to be 1 of the following produce/consume")
		os.Exit(1)
	}

	msgSizeString := (strings.Split(os.Args[2], "="))[1]
	msgSize, err := strconv.Atoi(msgSizeString)
	if err != nil || msgSize <= 0 {
		fmt.Println("msgSize has to be a positive number")
		os.Exit(1)
	}

	msgsCountString := (strings.Split(os.Args[3], "="))[1]
	msgsCount, err := strconv.Atoi(msgsCountString)
	if err != nil || msgsCount <= 0 {
		fmt.Println("msgCount has to be a positive number")
		os.Exit(1)
	}

	host := (strings.Split(os.Args[4], "="))[1]
	username := (strings.Split(os.Args[5], "="))[1]
	token := (strings.Split(os.Args[6], "="))[1]

	storageType := memphis.File
	storageTypeString := (strings.Split(os.Args[7], "="))[1]
	if storageTypeString == "file" {
		storageType = memphis.File
	} else if storageTypeString == "memory" {
		storageType = memphis.Memory
	} else {
		fmt.Println("storageType has to be 1 of the following file/memory")
		os.Exit(1)
	}

	replicasString := (strings.Split(os.Args[8], "="))[1]
	replicas, err := strconv.Atoi(replicasString)
	if err != nil || replicas <= 0 || replicas > 5 {
		fmt.Println("replicas has to be a positive number between 1-5")
		os.Exit(1)
	}

	cg := (strings.Split(os.Args[9], "="))[1]

	pullIntervalString := (strings.Split(os.Args[10], "="))[1]
	pullIntervalInt, err := strconv.Atoi(pullIntervalString)
	if err != nil || pullIntervalInt <= 0 {
		fmt.Println("pullInterval has to be a positive number")
		os.Exit(1)
	}
	pullInterval := time.Duration(pullIntervalInt) * time.Millisecond

	batchSizeString := (strings.Split(os.Args[11], "="))[1]
	batchSize, err := strconv.Atoi(batchSizeString)
	if err != nil || batchSize <= 0 {
		fmt.Println("batchSize has to be a positive number")
		os.Exit(1)
	}

	batchTTWString := (strings.Split(os.Args[12], "="))[1]
	batchTTWInt, err := strconv.Atoi(batchTTWString)
	if err != nil || batchTTWInt <= 0 {
		fmt.Println("batchTTW has to be a positive number")
		os.Exit(1)
	}
	batchTTW := time.Duration(batchTTWInt) * time.Millisecond

	c, err := memphis.Connect(host, username, token)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer c.Close()

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	s, err := c.CreateStation("station_"+timestamp,
		"benchmarks_factory",
		memphis.StorageTypeOpt(storageType),
		memphis.Replicas(replicas),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer s.Destroy()

	p, err := s.CreateProducer("prod_" + timestamp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	msg := getMsgInSize(msgSize)

	var start time.Time
	if opType == "produce" {
		start = time.Now()
	}

	for i := 0; i < msgsCount; i++ {
		p.Produce(msg)
	}

	if opType == "consume" {
		done := make(chan bool)
		cons, err := s.CreateConsumer("cons_"+timestamp,
			memphis.ConsumerGroup(cg),
			memphis.PullInterval(pullInterval),
			memphis.BatchSize(batchSize),
			memphis.BatchMaxWaitTime(batchTTW),
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		consumedMsgs := 0
		start = time.Now()

		cons.Consume(func(msgs []*memphis.Msg, err error) {
			if err == nil {
				for _, m := range msgs {
					m.Ack()
					consumedMsgs++
					if consumedMsgs >= msgsCount {
						done <- true
					}
				}
			}
		})
		_ = <-done
	}

	elapsed := time.Since(start).Seconds()
	msgsPerSec := float64(msgsCount) / float64(elapsed)
	mbPerSec := float64(msgSize*msgsCount) / float64(elapsed) / 1024 / 1024

	fmt.Printf("operation type: %s, msgs/sec: %v, MB/sec: %.2f", opType, math.Ceil(msgsPerSec), mbPerSec)
}
