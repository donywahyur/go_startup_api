# Use the Go 1.23.2 version
FROM golang:1.23.2-alpine as builder

# Set a flexible working directory
WORKDIR /home/gonuxt_startup/go

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire application
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image to serve the Go app
FROM alpine:latest

WORKDIR /home/gonuxt_startup/go

# Copy the compiled binary
COPY --from=builder /home/gonuxt_startup/go/main .

# Expose the API port
EXPOSE 8080

# Start the Go application
CMD ["./main"]