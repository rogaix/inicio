# Placeholder variable for the Go compiler
GOCMD=go

# The final binary to build
TARGET=server

# The image name for docker build
IMAGE_NAME=myapp

# The container name for docker run
CONTAINER_NAME=myapp_container

.PHONY: all build docker-build docker-run clean-docker-build

# This will run 'make build', 'make docker-build' and then 'make docker-run'
all: build docker-build docker-run

# This builds the binary and places it in the root dir of the project
build:
	$(GOCMD) build -v -o ./$(TARGET) ./cmd/...

# This builds the Docker image using the Dockerfile
docker-build:
	docker build -t $(IMAGE_NAME) .

# This runs the Docker container from the Docker image
docker-run:
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)
	docker run -p 8080:8080 --name $(CONTAINER_NAME) -d $(IMAGE_NAME)

# This will run 'make docker-build' and then remove the local binary
clean-docker-build:
	$(MAKE) docker-build
	rm $(TARGET)