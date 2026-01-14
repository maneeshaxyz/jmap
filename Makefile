.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -cover

testcov:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

covhtml: testcov
	go tool cover -html=coverage.out

lint: fmt vet test
	@echo "Code passed fmt, vet, and tests"

gci:
	golangci-lint run

#GOEXPERIMENT=jsonv2
build: lint
	go build ./cmd/jmapd/


dbuild:
	docker build -t jmap .

drun: dbuild
	docker run --rm --name jmap \
		-p 8443:8443 \
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
.PHONY: fmt vet test lint build run clean gci
