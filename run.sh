#!/bin/sh

#Benchmarks
#Produce, Replica 1, FILE
./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true

#Produce, Replica 3, FILE
./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true

#Produce, Replica 1, MEMORY
./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true

#Produce, Replica 3, MEMORY
./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true

#CONSUME, Replica 1, FILE
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done

#CONSUME, Replica 3, FILE
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done

#CONSUME, Replica 1, MEMORY
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done

#CONSUME, Replica 3, MEMORY
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done
for i in $(seq 10);do ./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=1 host=memphis.memphis.svc.cluster.local username=root pass=$PASS storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true ; sleep 0.1 ;done

