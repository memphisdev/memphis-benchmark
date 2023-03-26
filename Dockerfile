FROM --platform=$BUILDPLATFORM golang:1.19-alpine3.17 as build

WORKDIR $GOPATH/src/
COPY . .

ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-w" -a -o  .

RUN go install github.com/nats-io/natscli/nats@latest

FROM alpine:3.17
ENV GOPATH="/go/src"
WORKDIR /run

COPY --from=build $GOPATH/run.sh .
COPY --from=build /go/bin/* /usr/local/bin

RUN chmod +x ./run.sh
COPY --from=build $GOPATH/memphis-benchmarks .
