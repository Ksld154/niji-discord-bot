GOPATH = $(shell go env GOPATH)
export PATH := $(PATH):$(GOPATH)/bin

all: build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_parser pkg/nijiparser/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_bot cmd/main.go

clean: 
	rm -rf ./bin 