.DEFAULT_GOAL := build

VERSION=$(shell git describe --tags)

run:
	go run main.go

sum:
	go run main.go sum

BINARY_PATH=build/gitime

# use the -s and -w linker flags to strip the debugging information
LD_FLAGS_STRIP=-s -w

build: $(shell find . -name \*.go)
	# use the -s and -w linker flags to strip the debugging information
	go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH) .

build-windows-amd64: $(shell find . -name \*.go)
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH).exe .

release: build
	upx --ultra-brute build/gitime

test:
	go test `go list ./...`

coverage:
	go test `go list ./...` -coverprofile=coverage.txt -covermode=atomic

install: build
	sudo install build/gitime /usr/local/bin/

install-optimized: build release
	sudo install build/gitime /usr/local/bin/
