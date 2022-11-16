go run main.go opType=produce/consume/e2e msgSize=1050 msgCount=1000000 host=localhost username=root token=memphis storageType=file/memory replicas=1 cg=cg1 pullInterval=100 batchSize=100 batchTTW=500 concurrency=10 iterations=1 printHeaders=true asyncProduce=false deleteStations=false sleepMs=500 retentionType=default/msgs/bytes retentionValue=604800 

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
ITERATIONS
PRINT_HEADERS
ASYNC_PRODUCE
DELETE_STATIONS
SLEEP_MS
RETENTION_TYPE
RETENTION_VALUE

pullInterval/batchTTW - in microseconds
msgSize - bytes
concurrency - represents the number of go routines the program will create - max is 150
