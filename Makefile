.DEFAULT_GOAL := build

VERSION=$(shell git describe --tags)

run:
	go run main.go

sum:
	go run main.go sum

build: $(shell find . -name \*.go)
	go build -o build/gitime .

test:
	go test gitime/*.go

install: build run
	sudo install build/gitime /usr/local/bin/
