.PHONY: build

build:
	go build -o ./bin/updserver ./cmd/server 

lint:
	golint ./...

run: build
	./bin/updserver 8080
