package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/memphisdev/memphis.go"
)

type ExtConn struct {
	c    *memphis.Conn
	p    *memphis.Producer
	cons *memphis.Consumer
}

func getMsgInSize(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(1)
	}
	return bytes
}

func validateArgs() (string, int, int, string, string, string, memphis.StorageType, int, string, time.Duration, int, time.Duration, int, int, bool, bool, bool, int) {
	var opType, msgSizeString, msgsCountString, host, username, token, storageTypeString, replicasString, cg, pullIntervalString, batchSizeString,
		batchTTWString, concurrencyString, iterationsString,
		printHeadersString, asyncProduceString, deleteStationsString, sleepMsString string

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
		iterationsString = os.Getenv("ITERATIONS")
		printHeadersString = os.Getenv("PRINT_HEADERS")
		asyncProduceString = os.Getenv("ASYNC_PRODUCE")
		deleteStationsString = os.Getenv("DELETE_STATIONS")
		sleepMsString = os.Getenv("SLEEP_MS")
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
		iterationsString = (strings.Split(os.Args[14], "="))[1]
		printHeadersString = (strings.Split(os.Args[15], "="))[1]
		asyncProduceString = (strings.Split(os.Args[16], "="))[1]
		deleteStationsString = (strings.Split(os.Args[17], "="))[1]
		sleepMsString = (strings.Split(os.Args[18], "="))[1]
	}

	if opType != "produce" && opType != "consume" && opType != "e2e" {
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
	if err != nil || concurrency <= 0 {
		fmt.Println("concurrency has to be a positive number")
		os.Exit(1)
	}

	iterations, err := strconv.Atoi(iterationsString)
	if err != nil || iterations <= 0 {
		fmt.Println("iterations has to be a positive number")
		os.Exit(1)
	}

	printHeaders, err := strconv.ParseBool(printHeadersString)
	if err != nil {
		printHeaders = false
	}

	asyncProduce, err := strconv.ParseBool(asyncProduceString)
	if err != nil {
		asyncProduce = false
	}

	deleteStations, err := strconv.ParseBool(deleteStationsString)
	if err != nil {
		printHeaders = false
	}

	sleepMs, err := strconv.Atoi(sleepMsString)
	if err != nil || sleepMs <= 0 {
		fmt.Println("sleepMs has to be a positive number")
		os.Exit(1)
	}

	return opType, msgSize, msgsCount, host, username, token, storageType, replicas, cg, pullInterval, batchSize, batchTTW, concurrency, iterations, printHeaders, asyncProduce, deleteStations, sleepMs
}

func main() {
	opType, msgSize, msgsCount, host, username, token, storageType, replicas, cg, pullInterval, batchSize, batchTTW, concurrencyFactor, iterations, printHeaders, asyncProduce, deleteStations, sleepMs := validateArgs()
	msg := getMsgInSize(msgSize)
	if printHeaders {
		fmt.Println("operation,iterations,replica,msgSize,msgCount,pullInterval,batchSize,batchTTW,concurency,msgs/sec,MB/sec,time")
	}

	for i := 0; i < iterations; i++ {
		timestamp := strconv.Itoa(int(time.Now().Unix()))
		iterationsCount := strconv.Itoa(i)
		stationName := "station_" + timestamp + "_" + "_" + iterationsCount

		c, err := memphis.Connect(host, username, token)
		if err != nil {
			fmt.Println("Connect: " + err.Error())
			os.Exit(1)
		}
		s, err := c.CreateStation(stationName,
			memphis.StorageTypeOpt(storageType),
			memphis.Replicas(replicas),
		)
		if err != nil {
			fmt.Println("CreateStation: " + err.Error())
			os.Exit(1)
		}

		var extConn []*ExtConn
		for i := 0; i < concurrencyFactor; i++ {
			indexConcurrency := strconv.Itoa(i)
			c, err := memphis.Connect(host, username, token)
			if err != nil {
				fmt.Println("Connect: " + err.Error())
				os.Exit(1)
			}

			p, err := c.CreateProducer(stationName, "prod_"+indexConcurrency)
			if err != nil {
				fmt.Println("CreateProducer: " + err.Error())
				os.Exit(1)
			}

			ec := ExtConn{c: c, p: p}
			if opType == "e2e" || opType == "consume" {
				cons, err := c.CreateConsumer(stationName, "cons_"+indexConcurrency,
					memphis.ConsumerGroup(cg),
					memphis.PullInterval(pullInterval),
					memphis.BatchSize(batchSize),
					memphis.BatchMaxWaitTime(batchTTW),
					memphis.ConsumerErrorHandler(func(_ *memphis.Consumer, err error) {
						return
					}),
				)
				if err != nil {
					fmt.Println("CreateConsumer: " + err.Error())
					os.Exit(1)
				}

				ec.cons = cons
			}

			extConn = append(extConn, &ec)
		}

		// produce
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		if opType == "e2e" {
			wg.Add(concurrencyFactor)
		}

		for i := 0; i < concurrencyFactor; i++ {
			var count int
			if concurrencyFactor-1 == 0 { // run with single producer
				count = msgsCount
			} else if i == concurrencyFactor-1 {
				count = (msgsCount / concurrencyFactor) + (msgsCount % concurrencyFactor)
			} else {
				count = msgsCount / concurrencyFactor
			}

			go func(ec *ExtConn, msg []byte, count int, wg *sync.WaitGroup) {
				for i := 0; i < count; i++ {
					if asyncProduce {
						ec.p.Produce(msg, memphis.AsyncProduce())
					} else {
						ec.p.Produce(msg)
					}
				}
				wg.Done()
			}(extConn[i], msg, count, &wg)

			if opType == "e2e" {
				go func(ec *ExtConn, wg *sync.WaitGroup) {
					var wg2 sync.WaitGroup
					wg2.Add(1)
					done := false
					ec.cons.Consume(func(msgs []*memphis.Msg, err error) {
						if err == nil {
							for _, m := range msgs {
								m.Ack()
							}
						} else if !done {
							done = true
							wg2.Done()
						}
					})
					wg2.Wait()
					wg.Done()
				}(extConn[i], &wg)
			}
		}

		var start time.Time
		if opType == "produce" || opType == "e2e" {
			start = time.Now()
		}
		wg.Wait()

		// consume
		if opType == "consume" {
			var wg1 sync.WaitGroup
			wg1.Add(concurrencyFactor)

			for i := 0; i < concurrencyFactor; i++ {
				go func(ec *ExtConn, wg1 *sync.WaitGroup) {
					var wg2 sync.WaitGroup
					wg2.Add(1)
					done := false
					ec.cons.Consume(func(msgs []*memphis.Msg, err error) {
						if err == nil {
							for _, m := range msgs {
								m.Ack()
							}
						} else if !done {
							done = true
							wg2.Done()
						}
					})
					wg2.Wait()
					wg1.Done()
				}(extConn[i], &wg1)
			}

			start = time.Now()
			wg1.Wait()
		}

		elapsed := time.Since(start).Seconds()
		msgsPerSec := float64(msgsCount) / float64(elapsed)
		mbPerSec := float64(msgSize*msgsCount) / float64(elapsed) / 1024 / 1024

		for i := 0; i < concurrencyFactor; i++ {
			// extConn[i].p.Destroy()
			if opType == "consume" || opType == "e2e" {
				extConn[i].cons.Destroy()
			}
			extConn[i].c.Close()
		}
		if deleteStations {
			go s.Destroy()
		}

		fmt.Printf("%s,%v,%v,%v,%v,%v,%v,%v,%v,%f,%f,%f\n", opType, iterations, replicas, msgSize, msgsCount, pullInterval, batchSize, batchTTW, concurrencyFactor, math.Ceil(msgsPerSec), mbPerSec, float64(elapsed))
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}
}
