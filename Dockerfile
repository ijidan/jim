FROM golang:1.17

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /data/jim
COPY . .
RUN go build -o app

EXPOSE 8081 8082 8083
ENTRYPOINT ["/data/jim/app"]

