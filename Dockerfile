# Use the official Go image to create a build artifact.
FROM golang:1.25.0-bookworm AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod  ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /jmap .

FROM scratch

COPY --from=builder /jmap /jmap

EXPOSE 8080

ENTRYPOINT ["/jmap"]
