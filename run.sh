#!/bin/sh

#Benchmarks

./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true


./memphis-benchmarks opType=produce msgSize=512 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1024 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=256000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=produce msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true


./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=1 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true


./memphis-benchmarks opType=consume msgSize=512 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=true asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=1024 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=256000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true
./memphis-benchmarks opType=consume msgSize=1000000 produceRate=300 secondsToRun=10 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=memory replicas=3 pullInterval=3000 batchSize=3 batchTTW=10 concurrency=5 printHeaders=false asyncProduce=true deleteStations=true

