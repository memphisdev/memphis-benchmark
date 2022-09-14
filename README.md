go run main.go opType=produce/consume msgSize=1050 msgCount=1000000 host=localhost username=root token=memphis storageType=file/memory replicas=1 cg=cg1 pullInterval=100 batchSize=100 batchTTW=500 concurrency=10

pullInterval/batchTTW - in microseconds
msgSize - bytes