GOPATH = $(shell go env GOPATH)
export PATH := $(PATH):$(GOPATH)/bin

all: build

build:
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/niji_parser pkg/nijiparser/main.go
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/niji_scrapper pkg/nijionair/main.go

	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/bitly pkg/bitly/main.go
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/yt_picker pkg/ytpicker/main.go
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/help_msg pkg/helpmsg/main.go
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/utils pkg/utils/main.go
	CGO_ENABLED=0 env GOOS=linux go build -ldflags="-s -w" -o bin/tetrio pkg/tetrio/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/niji_bot cmd/main.go

clean: 
	rm -rf ./bin 