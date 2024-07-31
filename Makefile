# Placeholder variable for the Go compiler
GOCMD=go

# The final binary to build
TARGET=server

# The image name for docker build
IMAGE_NAME=inicio-app

# The container name for docker run
CONTAINER_NAME=inicio-app-1

# Define the directory for the Vue frontend
VUE_DIR=./web/vue

.PHONY: all build docker-build docker-run clean-docker-build frontend-build frontend-serve docker-compose-up frontend

# This will run 'make frontend-build', 'make build', 'make docker-build' and then 'make docker-run'
all: frontend-build build docker-build docker-run

# This builds the binary and places it in the root dir of the project
build:
	$(GOCMD) build -v -o ./$(TARGET) ./cmd/...

# This builds the Docker image using the Dockerfile
docker-build:
	docker build -t $(IMAGE_NAME) .

# This stops, removes any existing container, and then runs the Docker container using docker-compose
docker-run:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
	docker-compose up --build -d

# This will run 'make docker-build' and then remove the local binary
clean-docker-build:
	$(MAKE) docker-build
	rm $(TARGET)

# This builds the Vue frontend
frontend-build:
	cd $(VUE_DIR) && npm install && npm run build

# This serves the Vue frontend in development mode
frontend-serve:
	cd $(VUE_DIR) && npm install && npm run dev

# This copies the built frontend to the running Docker container
frontend: frontend-build
	@CONTAINER_NAME=$$(docker ps --filter "name=$(CONTAINER_NAME)" --format "{{.Names}}"); \
	if [ -z "$$CONTAINER_NAME" ]; then \
	    CONTAINER_NAME=$$(docker ps --filter "ancestor=$(IMAGE_NAME)" --format "{{.Names}}"); \
	fi; \
	if [ -z "$$CONTAINER_NAME" ]; then \
	    echo "Error: No running container found for image $(IMAGE_NAME) or name $(CONTAINER_NAME)"; \
	    exit 1; \
	fi; \
	docker cp $(VUE_DIR)/dist $$CONTAINER_NAME:/root/web/vue/
