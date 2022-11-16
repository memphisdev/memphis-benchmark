package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

func validateArgs() (string, int, int, int, string, string, string, memphis.StorageType, int, time.Duration, int, time.Duration, int, bool, bool, bool) {
	var opType, msgSizeString, produceRateString, secondsToRunString, host, username, token, storageTypeString, replicasString,
		pullIntervalString, batchSizeString, batchTTWString, concurrencyString, printHeadersString, asyncProduceString,
		deleteStationsString string

	if len(os.Args) < 2 {
		opType = os.Getenv("OP_TYPE")
		msgSizeString = os.Getenv("MSG_SIZE")
		produceRateString = os.Getenv("PRODUCE_RATE")
		secondsToRunString = os.Getenv("SECONDS_TO_RUN")
		host = os.Getenv("HOST")
		username = os.Getenv("USERNAME")
		token = os.Getenv("TOKEN")
		storageTypeString = os.Getenv("STORAGE_TYPE")
		replicasString = os.Getenv("REPLICAS")
		pullIntervalString = os.Getenv("PULL_INTERVAL")
		batchSizeString = os.Getenv("BATCH_SIZE")
		batchTTWString = os.Getenv("BATCH_TTW")
		concurrencyString = os.Getenv("CONCURRENCY")
		printHeadersString = os.Getenv("PRINT_HEADERS")
		asyncProduceString = os.Getenv("ASYNC_PRODUCE")
		deleteStationsString = os.Getenv("DELETE_STATIONS")
	} else {
		opType = (strings.Split(os.Args[1], "="))[1]
		msgSizeString = (strings.Split(os.Args[2], "="))[1]
		produceRateString = (strings.Split(os.Args[3], "="))[1]
		secondsToRunString = (strings.Split(os.Args[4], "="))[1]
		host = (strings.Split(os.Args[5], "="))[1]
		username = (strings.Split(os.Args[6], "="))[1]
		token = (strings.Split(os.Args[7], "="))[1]
		storageTypeString = (strings.Split(os.Args[8], "="))[1]
		replicasString = (strings.Split(os.Args[9], "="))[1]
		pullIntervalString = (strings.Split(os.Args[10], "="))[1]
		batchSizeString = (strings.Split(os.Args[11], "="))[1]
		batchTTWString = (strings.Split(os.Args[12], "="))[1]
		concurrencyString = (strings.Split(os.Args[13], "="))[1]
		printHeadersString = (strings.Split(os.Args[14], "="))[1]
		asyncProduceString = (strings.Split(os.Args[15], "="))[1]
		deleteStationsString = (strings.Split(os.Args[16], "="))[1]
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

	produceRate, err := strconv.Atoi(produceRateString)
	if err != nil || produceRate <= 0 {
		fmt.Println("produceRate has to be a positive number")
		os.Exit(1)
	}

	secondsToRun, err := strconv.Atoi(secondsToRunString)
	if err != nil || secondsToRun <= 0 {
		fmt.Println("secondsToRun has to be a positive number")
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

	return opType, msgSize, produceRate, secondsToRun, host, username, token, storageType, replicas, pullInterval, batchSize, batchTTW, concurrency, printHeaders, asyncProduce, deleteStations
}

func main() {
	opType, msgSize, produceRate, secondsToRun, host, username, token, storageType, replicas, pullInterval, batchSize, batchTTW, concurrencyFactor, printHeaders, asyncProduce, deleteStations := validateArgs()

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	stationName := "station_" + timestamp
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

	// initialize connection resources
	var extConn []*ExtConn
	for j := 0; j < concurrencyFactor; j++ {
		index1 := strconv.Itoa(j)
		c, err := memphis.Connect(host, username, token)
		if err != nil {
			fmt.Println("Connect: " + err.Error())
			os.Exit(1)
		}
		p, err := c.CreateProducer(stationName, "prod_"+index1)
		if err != nil {
			fmt.Println("CreateProducer: " + err.Error())
			os.Exit(1)
		}
		ec := ExtConn{c: c, p: p}
		if opType == "e2e" || opType == "consume" {
			cons, err := c.CreateConsumer(stationName, "cons_"+index1,
				memphis.ConsumerGroup("group1"),
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

	msg := getMsgInSize(msgSize)
	if printHeaders {
		fmt.Println("operation,msgSize,produceRate,storageType,replicas,pullInterval,batchSize,batchTTW,concurency,msgsCount,latency")
	}

	for i := 0; i < secondsToRun; i++ {
		finish := make(chan bool)

		// produce/e2e
		for i := 0; i < concurrencyFactor; i++ {
			var count int
			if concurrencyFactor-1 == 0 { // run with single producer
				count = produceRate
			} else if i == concurrencyFactor-1 {
				count = (produceRate / concurrencyFactor) + (produceRate % concurrencyFactor)
			} else {
				count = produceRate / concurrencyFactor
			}

			go func(ec *ExtConn, msg []byte, count int, ch chan bool) {
				for i := 0; i < count; i++ {
					if asyncProduce {
						ec.p.Produce(msg, memphis.AsyncProduce())
					} else {
						ec.p.Produce(msg)
					}
				}
				ch <- true
			}(extConn[i], msg, count, finish)

			if opType == "e2e" {
				go func(ec *ExtConn, ch chan bool) {
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
					ch <- true
				}(extConn[i], finish)
			}
		}
		var start time.Time
		if opType == "produce" || opType == "e2e" {
			start = time.Now()
		}

		// consume
		if opType == "consume" {
			for i := 0; i < concurrencyFactor; i++ {
				go func(ec *ExtConn, ch chan bool) {
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
					ch <- true
				}(extConn[i], finish)
			}

			start = time.Now()
		}

		var latency, msgsCount int64
		count := 0
		done := false
		for {
			if done {
				break
			} else {
				select {
				case <-finish:
					count++
					if (opType == "produce" && count == concurrencyFactor) || (opType == "e2e" && count == 2*concurrencyFactor) || (opType == "consume" && count == 2*concurrencyFactor) {
						latency = time.Since(start).Milliseconds()
					}

				case <-time.After(time.Second * 1):
					command := fmt.Sprintf("nats stream info %s --server=%s:6666 --user=%s", stationName, host, token)
					cmd := exec.Command("bash", "-c", command)
					var outb bytes.Buffer
					cmd.Stdout = &outb
					err = cmd.Run()
					if err != nil {
						cmd = exec.Command("sh", "-c", command)
						cmd.Stdout = &outb
						err = cmd.Run()
						if err != nil {
							fmt.Println("info station: " + err.Error())
							os.Exit(1)
						}
					}
					cmdOut := outb.String()
					cmdOut = strings.Split(cmdOut, "  Messages: ")[1]
					cmdOut = strings.Split(cmdOut, "\n")[0]
					cmdOut = strings.Replace(cmdOut, ",", "", -1)
					num, _ := strconv.Atoi(cmdOut)
					msgsCount = int64(num)

					if latency == 0 {
						latency = 1000
					}
					fmt.Printf("%s,%v,%v,%s,%v,%v,%v,%v,%v,%v,%v\n", opType, msgSize, produceRate, storageType, replicas, pullInterval, batchSize, batchTTW, concurrencyFactor, msgsCount, latency)

					if opType == "consume" || opType == "e2e" {
						time.Sleep(10 * time.Second)
						command = fmt.Sprintf("nats stream purge %s -f --server=%s:6666 --user=%s", stationName, host, token)
						cmd = exec.Command("bash", "-c", command)
						err = cmd.Run()
						if err != nil {
							cmd = exec.Command("sh", "-c", command)
							err = cmd.Run()
							if err != nil {
								fmt.Println("purge station: " + err.Error())
								os.Exit(1)
							}
						}
					}

					done = true
				}
			}
		}
	}

	if deleteStations {
		go s.Destroy()
	}
}
