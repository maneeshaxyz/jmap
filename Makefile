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
	go build .

linux:
	GOOS=linux GOARCH=amd64 go build -o jmap

run: build
	./jmap

clean:
	rm -f jmap

# --- Phony targets ---
.PHONY: fmt vet test lint build run clean
