.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint: fmt vet test
	@echo "Code passed fmt, vet, and tests"

build: lint
	go build ./cmd/jmapd/
#GOEXPERIMENT=jsonv2

dbuild:
	docker build -t jmap .

drun: dbuild
	docker run --rm --name jmap \
		-p 8080:8080 \
		-v "$(PWD)/server.crt:/certs/server.crt:ro" \
		-v "$(PWD)/server.key:/certs/server.key:ro" \
		jmap

linux:
	GOOS=linux GOARCH=amd64 go build -o jmap

run: build
	./jmapd

clean:
	rm -f jmap

# --- Phony targets ---
.PHONY: fmt vet test lint build run clean
