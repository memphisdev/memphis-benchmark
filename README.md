go run main.go opType=produce/e2e msgSize=1050 produceRate=1000000 secondsToRun=5 host=localhost username=root token=memphis storageType=file/memory replicas=1 pullInterval=100 batchSize=100 batchTTW=500 concurrency=10 printHeaders=true asyncProduce=false deleteStations=false

or go run main.go with the following environment variables
OP_TYPE
MSG_SIZE
PRODUCE_RATE (msgs/sec) // MSG_COUNT
SECONDS_TO_RUN // ITERATIONS
HOST
USERNAME
TOKEN
STORAGE_TYPE
REPLICAS
PULL_INTERVAL
BATCH_SIZE
BATCH_TTW
CONCURRENCY
PRINT_HEADERS
ASYNC_PRODUCE
DELETE_STATIONS

pullInterval/batchTTW - in microseconds
msgSize - bytes
concurrency - represents the number of go routines the program will create - max is 150
