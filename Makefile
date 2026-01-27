.PHONY: build run

build:
	mkdir -p bin
	go build -o bin/monkey ./cmd/monkey/main.go

run:
	go run ./cmd/monkey/main.go
