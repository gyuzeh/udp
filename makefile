.PHONY: build

build:
	go build -o ./bin/updserver ./cmd/server 
	go build -o ./bin/udpClient ./cmd/client 

lint:
	golint ./...

server: build
	./bin/updserver 8080

client: build
	./bin/udpClient 127.0.0.1 8080
