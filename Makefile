ROOT_DIR := $(shell basename $(CURDIR))

build: format
	GO111MODULE=off GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/${ROOT_DIR}-linux-amd64
	GO111MODULE=off GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/${ROOT_DIR}-darwin-amd64

format:
	GO111MODULE=off go fmt ./...
