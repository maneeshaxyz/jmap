.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint: fmt vet test
	@echo "Code passed fmt, vet, and tests ✅"

build: lint
	go build -o server main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o server

run: build
	./server

clean:
	rm -f server


# --- Phony targets ---
.PHONY: fmt vet test lint build run clean
