package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

func validateArgs() (string, int, int, string, string, string, memphis.StorageType, int, string, time.Duration, int, time.Duration, int) {
	var opType, msgSizeString, msgsCountString, host, username, token, storageTypeString, replicasString, cg, pullIntervalString, batchSizeString, batchTTWString, concurrencyString string

	if len(os.Args) < 2 {
		opType = os.Getenv("OP_TYPE")
		msgSizeString = os.Getenv("MSG_SIZE")
		msgsCountString = os.Getenv("MSG_COUNT")
		host = os.Getenv("HOST")
		username = os.Getenv("USERNAME")
		token = os.Getenv("TOKEN")
		storageTypeString = os.Getenv("STORAGE_TYPE")
		replicasString = os.Getenv("REPLICAS")
		cg = os.Getenv("CG")
		pullIntervalString = os.Getenv("PULL_INTERVAL")
		batchSizeString = os.Getenv("BATCH_SIZE")
		batchTTWString = os.Getenv("BATCH_TTW")
		concurrencyString = os.Getenv("CONCURRENCY")
	} else {
		opType = (strings.Split(os.Args[1], "="))[1]
		msgSizeString = (strings.Split(os.Args[2], "="))[1]
		msgsCountString = (strings.Split(os.Args[3], "="))[1]
		host = (strings.Split(os.Args[4], "="))[1]
		username = (strings.Split(os.Args[5], "="))[1]
		token = (strings.Split(os.Args[6], "="))[1]
		storageTypeString = (strings.Split(os.Args[7], "="))[1]
		replicasString = (strings.Split(os.Args[8], "="))[1]
		cg = (strings.Split(os.Args[9], "="))[1]
		pullIntervalString = (strings.Split(os.Args[10], "="))[1]
		batchSizeString = (strings.Split(os.Args[11], "="))[1]
		batchTTWString = (strings.Split(os.Args[12], "="))[1]
		concurrencyString = (strings.Split(os.Args[13], "="))[1]
	}

	if opType != "produce" && opType != "consume" {
		fmt.Println("opType has to be 1 of the following produce/consume")
		os.Exit(1)
	}

	msgSize, err := strconv.Atoi(msgSizeString)
	if err != nil || msgSize <= 0 {
		fmt.Println("msgSize has to be a positive number")
		os.Exit(1)
	}

	msgsCount, err := strconv.Atoi(msgsCountString)
	if err != nil || msgsCount <= 0 {
		fmt.Println("msgCount has to be a positive number")
		os.Exit(1)
	}

	storageType := memphis.File
	if storageTypeString == "file" {
		storageType = memphis.File
	} else if storageTypeString == "memory" {
		storageType = memphis.Memory
	} else {
		fmt.Println("storageType has to be 1 of the following file/memory")
		os.Exit(1)
	}

	replicas, err := strconv.Atoi(replicasString)
	if err != nil || replicas <= 0 || replicas > 5 {
		fmt.Println("replicas has to be a positive number between 1-5")
		os.Exit(1)
	}

	pullIntervalInt, err := strconv.Atoi(pullIntervalString)
	if err != nil || pullIntervalInt <= 0 {
		fmt.Println("pullInterval has to be a positive number")
		os.Exit(1)
	}
	pullInterval := time.Duration(pullIntervalInt) * time.Microsecond

	batchSize, err := strconv.Atoi(batchSizeString)
	if err != nil || batchSize <= 0 {
		fmt.Println("batchSize has to be a positive number")
		os.Exit(1)
	}

	batchTTWInt, err := strconv.Atoi(batchTTWString)
	if err != nil || batchTTWInt <= 0 {
		fmt.Println("batchTTW has to be a positive number")
		os.Exit(1)
	}
	batchTTW := time.Duration(batchTTWInt) * time.Microsecond

	concurrency, err := strconv.Atoi(concurrencyString)
	if err != nil || batchSize <= 0 {
		fmt.Println("concurrency has to be a positive number")
		os.Exit(1)
	}

	return opType, msgSize, msgsCount, host, username, token, storageType, replicas, cg, pullInterval, batchSize, batchTTW, concurrency
}

func main() {
	opType, msgSize, msgsCount, host, username, token, storageType, replicas, cg, pullInterval, batchSize, batchTTW, concurrency := validateArgs()

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

	// produce
	msg := getMsgInSize(msgSize)

	concurrencyFactor := runtime.NumCPU()
	if concurrency < concurrencyFactor {
		concurrencyFactor = concurrency
	}

	var producers []*memphis.Producer
	for i := 0; i < concurrencyFactor; i++ {
		index := strconv.Itoa(i)
		p, err := s.CreateProducer("prod_" + index)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		producers = append(producers, p)
	}

	var start time.Time
	if opType == "produce" {
		start = time.Now()
	}

	var wg sync.WaitGroup
	wg.Add(concurrencyFactor)
	for i := 0; i < concurrencyFactor; i++ {
		var count int
		if concurrencyFactor-1 == 0 { // run with single producer
			count = msgsCount
		} else if i == concurrencyFactor-1 {
			count = (msgsCount / concurrencyFactor) + (msgsCount % concurrencyFactor)
		} else {
			count = msgsCount / concurrencyFactor
		}

		go func(p *memphis.Producer, msg []byte, count int, wg *sync.WaitGroup) {
			for i := 0; i < count; i++ {
				p.Produce(msg)
			}
			wg.Done()
		}(producers[i], msg, count, &wg)
	}
	wg.Wait()

	// consume
	if opType == "consume" {
		var consumers []*memphis.Consumer
		for i := 0; i < concurrencyFactor; i++ {
			index := strconv.Itoa(i)
			cons, err := s.CreateConsumer("cons_"+index,
				memphis.ConsumerGroup(cg),
				memphis.PullInterval(pullInterval),
				memphis.BatchSize(batchSize),
				memphis.BatchMaxWaitTime(batchTTW),
			)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			consumers = append(consumers, cons)
		}

		start = time.Now()

		var wg1 sync.WaitGroup
		wg1.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			go func(c *memphis.Consumer, wg *sync.WaitGroup) {
				quit := make(chan bool)
				c.Consume(func(msgs []*memphis.Msg, err error) {
					if err == nil {
						for _, m := range msgs {
							m.Ack()
						}
					} else {
						quit <- true
					}
				})
				<-quit
				wg.Done()
				return
			}(consumers[i], &wg1)
		}
		wg1.Wait()
	}

	elapsed := time.Since(start).Seconds()
	msgsPerSec := float64(msgsCount) / float64(elapsed)
	mbPerSec := float64(msgSize*msgsCount) / float64(elapsed) / 1024 / 1024

	fmt.Printf("operation type: %s, msgs/sec: %v, MB/sec: %.2f total time: %.2f: ", opType, math.Ceil(msgsPerSec), mbPerSec, float64(elapsed))
}
