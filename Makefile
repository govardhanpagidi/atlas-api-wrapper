.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
#goos=linux
#goarch=amd64
# Variables
DOCKER_COMPOSE = docker-compose
APP_NAME = atlas_api_helper
APP_IMAGE = $(APP_NAME)-image
APP_PORT = 8080

goos=darwin
goarch=amd64

DEBUG_FLAGS=all=-N -l
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
LINKER_FLAGS=-s -w -X github.com/govardhanpagidi/atlas-api-helper/util.defaultLogLevel=debug -X github.com/govardhanpagidi/atlas-api-helper/version.Version=${CFNREP_GIT_SHA}

build:
	@echo "==> Building handler binary"
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o atlas_api_helper ./main.go

debug:
	@echo "==> Building handler binary for debugging"
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -gcflags="$(DEBUG_FLAGS)" -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/handler cmd/main.go

swaggers:
	swag init
	mv ./docs/swagger.json ./docs/doc.json && mv ./docs/swagger.yaml ./tools/openapi/swagger.yaml  \
	&& mv ./docs/doc.json ./tools/openapi/swagger.json 

clean:
	rm -rf bin

# Makefile for Go application with Docker Compose


.PHONY: build run stop clean

# Build and start the application using Docker Compose
run: 
	$(DOCKER_COMPOSE) up -d

# Stop and remove the Docker containers
stop:
	$(DOCKER_COMPOSE) down

# Clean up the built Go binary
clean:
	rm -f $(APP_NAME)

# Run the application in development mode (e.g., without Docker Compose)
dev: build
	./$(APP_NAME)


.PHONY: openapi
openapi: swaggers
	echo "Running OpenAPI Generation and Validation process"
	$(MAKE) -C tools clean_client
	echo "Running client generation"
	$(MAKE) -C tools generate_client
	echo "Running OpenAPI Validation"
	$(MAKE) -C tools generate_tests

default: run
