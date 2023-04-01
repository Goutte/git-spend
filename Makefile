# Welcome to the Makefile, curious friend.
# Run any recipe with `make <recipe>`.
# Improve these recipes at will, I'm learning as I go and they feel not right sometimes.

.DEFAULT_GOAL := build

# We have a directory named "build" and another named "test", so…
.PHONY: build test

VERSION=$(shell git describe --tags)

BINARY_PATH=build/git-spend
SOURCE=.
# /!. CAREFUL: THIS DIRECTORY WILL BE RM -RF 'ed
TMP_MAN_PATH=/tmp/git-spend-man

# Use the -s and -w linker flags to strip the debugging information
LD_FLAGS_STRIP=-s -w


depend:
	go get

run:
	go run .

sum:
	go run . sum

clean:
	rm -f "$(BINARY_PATH)"
	rm -f "$(BINARY_PATH).upx"
	rm -f "$(BINARY_PATH)-coverage"
	rm -f "$(BINARY_PATH).exe"
	rm -f "$(BINARY_PATH).000"
	rm -f test-coverage/*
	rm -rf "$(TMP_MAN_PATH)"

build:# $(shell find . -name \*.go)
	go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH) $(SOURCE)

build-coverage:
	go build -cover -o $(BINARY_PATH)-coverage $(SOURCE)

build-linux-arm64: $(shell find . -name \*.go)
	GOOS=windows GOARCH=arm64 go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH).arm64 $(SOURCE)

build-windows-amd64: $(shell find . -name \*.go)
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS_STRIP)" -o $(BINARY_PATH).exe $(SOURCE)

release: clean build build-windows-amd64 build-linux-arm64
	upx --ultra-brute $(BINARY_PATH)
	upx --ultra-brute $(BINARY_PATH).exe

test: test-unit

test-all: test-unit test-acceptance

test-unit:
	go test `go list ./...`

test-acceptance: build test-acceptance-depend
	test/bats/bin/bats test

test-acceptance-depend:
	git submodule update --init --recursive

coverage:
	go test `go list ./...` -coverprofile=coverage-unit.txt -covermode=atomic

coverage-acceptance: clean build-coverage
	GIT_SPEND_COVERAGE=1 test/bats/bin/bats test
	go tool covdata textfmt -i=test-coverage/ -o coverage-integration.txt

install: build
	sudo install "$(BINARY_PATH)" /usr/local/bin/

install-release: release
	sudo install "$(BINARY_PATH)" /usr/local/bin/

man: clean
	mkdir -p "$(TMP_MAN_PATH)"
	go run . man --output "$(TMP_MAN_PATH)"
	echo "man pages were generated in $(TMP_MAN_PATH)"

install-man: build
	#sudo go run . man --install  # nope, `go` may not be available to `root`
	sudo "$(BINARY_PATH)" man --install
	# … same as
	#sudo mkdir -p /usr/local/share/man/man8
	#sudo install "$(TMP_MAN_PATH)"/git-spend*.8 /usr/local/share/man/man8/
