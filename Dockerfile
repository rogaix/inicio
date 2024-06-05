# Start from golang base image
FROM golang:alpine as builder

# Setting up workspace
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code from the current directory to the image
COPY . .

# Build the Go app
RUN go build -o server ./cmd/...

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

# Copy the binary from builder
COPY --from=builder /app/server .

# Expose port to the Docker host
EXPOSE 8080

# Command to run the binary
CMD ["./server"]