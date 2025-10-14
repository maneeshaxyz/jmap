.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -o jmap

run: build
	./jmap

format:
	go fmt

dbuild:
	docker build -t jmap .

drun:
	docker run jmap

.PHONY: fmt vet build run
