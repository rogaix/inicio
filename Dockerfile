# Build the Go app
FROM golang:alpine as builder

# Setting up workspace
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

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

# Install necessary packages including dcron
RUN apk --no-cache add ca-certificates tzdata bash curl mysql-client dcron

# Set timezone data for MySQL
ENV MYSQL_TZINFO=/usr/share/zoneinfo

WORKDIR /root

# Copy the binary from the builder
COPY --from=builder /app/server .

# Copy the built Vue files from the frontend-builder
COPY --from=frontend-builder /app/dist ./web/vue/dist

# Create a crontab file and copy it to the appropriate location
COPY crontab /etc/crontabs/root

# Make the entrypoint script executable
COPY entrypoint.sh /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh

# Expose port to the Docker host
EXPOSE 8080

# Set environment variables for MySQL
ENV MYSQL_HOST=${MYSQL_HOST}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}

# Set the entrypoint to the script
ENTRYPOINT ["/root/entrypoint.sh"]
