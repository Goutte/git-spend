.DEFAULT_GOAL := build

VERSION=$(shell git describe --tags)

run:
	go run main.go

sum:
	go run main.go sum

build: $(shell find . -name \*.go)
	# use the -s and -w linker flags to strip the debugging information
	go build -ldflags="-s -w" -o build/gitime .

release: build
	upx --brute build/gitime

test:
	go test gitime/*.go

install: release
	sudo install build/gitime /usr/local/bin/
