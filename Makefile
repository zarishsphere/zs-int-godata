.PHONY: build run test

build:
	go build -o tmp/server cmd/server/main.go

run: build
	./tmp/server

test:
	go test ./... -v
