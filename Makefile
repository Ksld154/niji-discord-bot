GOPATH = $(shell go env GOPATH)
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: all build test clean

all: build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_parser pkg/nijiparser/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_scrapper pkg/nijionair/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/bitly pkg/bitly/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/yt_picker pkg/ytpicker/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/help_msg pkg/helpmsg/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/utils pkg/utils/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_bot cmd/main.go

test:
	# go test -v -cover=true pkg/nijiparser/main_test.go pkg/nijiparser/main.go
	
	go test -short pkg/nijiparser/main_test.go pkg/nijiparser/main.go
	go test -short -cover=true pkg/bitly/main_test.go pkg/bitly/main.go
	go test -short -cover=true pkg/ytpicker/main_test.go pkg/ytpicker/main.go
	go test -short -cover=true pkg/utils/main_test.go pkg/utils/main.go


clean: 
	rm -rf ./bin 