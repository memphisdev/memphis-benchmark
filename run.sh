#!/bin/sh

#Benchmarks

./memphis-benchmarks opType=produce msgSize=256 msgCount=100000 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 cg=cg1 pullInterval=100 batchSize=20 batchTTW=40 concurrency=150 iterations=10 stationsCount=1 printHeader=false >> 256.csv
./memphis-benchmarks opType=produce msgSize=512 msgCount=100000 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 cg=cg1 pullInterval=100 batchSize=20 batchTTW=40 concurrency=150 iterations=10 stationsCount=1 printHeader=false >> 512.csv
./memphis-benchmarks opType=produce msgSize=1024 msgCount=100000 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 cg=cg1 pullInterval=100 batchSize=20 batchTTW=40 concurrency=150 iterations=10 stationsCount=1 printHeader=false >> 1024.csv
./memphis-benchmarks opType=produce msgSize=5120 msgCount=100000 host=memphis-cluster.memphis.svc.cluster.local username=root token=$TOKEN storageType=file replicas=1 cg=cg1 pullInterval=100 batchSize=20 batchTTW=40 concurrency=150 iterations=10 stationsCount=1 printHeader=false >> 5120.csv
