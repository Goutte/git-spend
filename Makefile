.DEFAULT_GOAL := build

run:
	go run main.go

build: $(shell find . -name \*.go)
	go build -o build/gitime main.go

test:
	go test gitime/*.go

install: build run
	sudo install build/gitime /usr/local/bin/
