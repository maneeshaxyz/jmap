FROM golang:1.25-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /jmap ./cmd/jmapd

FROM scratch

COPY --from=builder /jmap /jmap

EXPOSE 8443

ENTRYPOINT ["/jmap"]
