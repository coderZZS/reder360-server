.PHONY: all build run run-main gotool clean help

BINARY="gin-cli"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@fresh

run-main:
	@go run ./main.go

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码，并编译成二进制文件"
	@echo "make build - 编译 Go 代码，并生成二进制文件"
	@echo "make run - 直接运行Go代码，并保持热更新"
	@echo "make run-main - 直接运行Go代码，没有热更新"
	@echo "make clean - 移除二进制文件和vim swap files"