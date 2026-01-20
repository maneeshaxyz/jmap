FROM golang:1.25-trixie AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /jmap ./cmd/jmapd

FROM scratch

COPY --from=builder /jmap /jmap
COPY --from=builder /app/server.crt /server.crt
COPY --from=builder /app/server.key /server.key

EXPOSE 8443

ENTRYPOINT ["/jmap"]
