BUILD_NAME := tfdocs-format-template

BUILD_DIR     := bin
PLUGIN_FOLDER ?= ~/.tfdocs.d/plugins

GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

CUR_VERSION ?=v1.0.0
DEFAULT_TAG  ?= $(shell echo "$(CUR_VERSION)" | tr -d 'v')
DOCKER_REGISTRY :=quay.io
PROJECT_OWNER := terraform-focs
DOCKER_IMAGE := $(DOCKER_REGISTRY)/$(PROJECT_OWNER)/$(BUILD_NAME)
DOCKER_TAG   ?= $(DEFAULT_TAG)

all: clean verify build

clean:
	rm -rf ./$(BUILD_DIR) ./${BUILD_NAME}

verify:
	go mod verify

build: clean
	CGO_ENABLED=0 go build -o ./$(BUILD_DIR)/$(GOOS)-$(GOARCH)/$(BUILD_NAME)

install: build
	mkdir -p $(PLUGIN_FOLDER)
	mv ./$(BUILD_DIR)/$(GOOS)-$(GOARCH)/$(BUILD_NAME) $(PLUGIN_FOLDER)

docker: ## Build Docker image
	docker build --pull --tag $(DOCKER_IMAGE):$(DOCKER_TAG) --file Dockerfile .

push: ## Push Docker image
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: all clean verify build install help