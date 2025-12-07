# Build stage
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go build -o student-ci-challenge

# Run stage
FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/student-ci-challenge /app/student-ci-challenge
EXPOSE 8080
ENTRYPOINT ["/app/student-ci-challenge"]
