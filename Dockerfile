# Use the official Go image to create a build artifact.
FROM golang:1.25.0-bookworm AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /jmap ./cmd/jmapd

FROM scratch

COPY --from=builder /jmap /jmap

EXPOSE 8080

ENTRYPOINT ["/jmap"]
