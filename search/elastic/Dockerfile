# Use multi-stage build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.* ./
RUN go mod download

# Copy the source code
COPY main.go ./

# Build the application
RUN go build -o main

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]