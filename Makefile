# Welcome to the Makefile, curious friend.
# Run any recipe with `make <recipe>`.
# Improve these recipes at will, I'm learning as I go and they feel not right sometimes.

.DEFAULT_GOAL := build

# We have a directory named "build", so…
.PHONY: build

VERSION=$(shell git describe --tags)

run:
	go run main.go

sum:
	go run main.go sum

BINARY_PATH=build/gitime

# use the -s and -w linker flags to strip the debugging information
LD_FLAGS_STRIP=-s -w

clean:
	rm -f $(BINARY_PATH)

build:# $(shell find . -name \*.go)
	# use the -s and -w linker flags to strip the debugging information
	go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH) .

build-windows-amd64:# $(shell find . -name \*.go)
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH).exe .

release: clean build build-windows-amd64
	upx --ultra-brute build/gitime

test:
	go test `go list ./...`

coverage:
	go test `go list ./...` -coverprofile=coverage.txt -covermode=atomic

install: build
	sudo install build/gitime /usr/local/bin/

install-release: release
	sudo install build/gitime /usr/local/bin/
