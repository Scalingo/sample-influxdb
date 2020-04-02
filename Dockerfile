FROM golang:1.14
MAINTAINER Étienne Michon "etienne@scalingo.com"

RUN go get github.com/cespare/reflex

WORKDIR $GOPATH/src/github.com/Scalingo/sample-influxdb
COPY . ./

RUN go build

EXPOSE 8086

CMD ["./sample-influxdb"]
