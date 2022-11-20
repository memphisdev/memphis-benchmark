FROM golang:1.18-alpine3.15 as build
WORKDIR $GOPATH/src/
COPY . .
RUN CGO_ENABLED=0 go build  -ldflags="-s -w" -a -o .
RUN go install github.com/nats-io/natscli/nats@latest

FROM alpine:3.15
ENV GOPATH="/go/src"
WORKDIR /run

COPY --from=build $GOPATH/run.sh .
COPY --from=build /go/bin/* /usr/local/bin

RUN chmod +x ./run.sh
COPY --from=build $GOPATH/memphis-benchmarks .
