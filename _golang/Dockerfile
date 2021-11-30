FROM golang:1.17

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

ARG WORK_DIR="/var/www/jim"

RUN apt update \
    apt-get install net-tools \
    apt-get install iputils-ping \
    apt-get install iproute2

WORKDIR $WORK_DIR
EXPOSE 8081 8082 8083

RUN go build "cmd/main/main.go"

ENTRYPOINT ["./main"]

