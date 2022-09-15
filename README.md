go run main.go opType=produce/consume msgSize=1050 msgCount=1000000 host=localhost username=root token=memphis storageType=file/memory replicas=1 cg=cg1 pullInterval=100 batchSize=100 batchTTW=500 concurrency=10

or go run main.go with the following environment variables
OP_TYPE
MSG_SIZE
MSG_COUNT
HOST
USERNAME
TOKEN
STORAGE_TYPE
REPLICAS
CG
PULL_INTERVAL
BATCH_SIZE
BATCH_TTW
CONCURRENCY

pullInterval/batchTTW - in microseconds
msgSize - bytes
concurrency - represents the number of go routines the program will create - max is 150
