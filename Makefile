
default: build

run:
	go run main.go

build:
	go build -o build/gitime main.go

test:
	go test gitime/*.go
