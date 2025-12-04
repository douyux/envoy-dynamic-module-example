.PHONY: build docker-build push clean help

# Variables
MODULE_NAME := libgo_module.so
COMMIT_ID := $(shell git rev-parse --short HEAD)
IMAGE_TAG := v1.36.2-go-$(COMMIT_ID)
IMAGE_FULL_NAME := registry.qingteng.cn/bxwu/envoy/envoy-examples

PLATFORMS := linux/amd64,linux/arm64

build:
	@echo "Building $(MODULE_NAME) for current platform..."
	CGO_ENABLED=1 go build -buildmode=c-shared -o $(MODULE_NAME) ./examples
	@echo "Build complete: $(MODULE_NAME)"

docker-build:
	@echo "Building Docker image $(IMAGE_FULL_NAME):$(IMAGE_TAG) for local platform..."
	docker build -t $(IMAGE_FULL_NAME):v1.36.2-go-latest .
	@echo "Docker image built successfully: $(IMAGE_FULL_NAME):v1.36.2-go-latest"
