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

# Build the Vue app
FROM node:alpine as frontend-builder

WORKDIR /app

# Copy the Vue project files
COPY web/vue/package*.json ./
RUN npm install
COPY web/vue .

# Build the Vue app
RUN npm run build

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

# Copy the binary from the builder
COPY --from=builder /app/server .

# Copy the built Vue files from the frontend-builder
COPY --from=frontend-builder /app/dist ./web/vue/dist

# Expose port to the Docker host
EXPOSE 8080

# Command to run the binary
CMD ["./server"]
