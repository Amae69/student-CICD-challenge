# Build stage
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go build -o app

# Run stage
FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
