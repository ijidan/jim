APP = jim
PACKAGE =
OUTPUT_BUILD_DIR = /data/jim


.PHONY : mod-down build run docker-compose grpc-server grpc-client clean help
help:
	@echo "make mod-down -  Go Mod下载"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make docker-compose - docker-compose直接运行 Go 代码"
	@echo "make grpc-server - 运行grpc server代码"
	@echo "make grpc-client - 运行grpc client代码"

	@echo "make clean - 清除vendor"
mod-down:
	@go mod download
build:
	@echo "Building  app..."
	@mkdir -p $(OUTPUT_BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o $(OUTPUT_BUILD_DIR)/$(APP) -ldflags '-w -s'
run: build
	@go run $(OUTPUT_BUILD_DIR)/$(APP)
docker-compose:
	@echo "Building $(app) app in docker..."
	@docker-compose up -d
grpc-server: mod-down
	@go run cmd/grpc/server/main.go
grpc-client:mod-down
	@go run cmd/grpc/client/main.go
clean:
	@echo "Cleaning..."
	@rm -rf vendor
