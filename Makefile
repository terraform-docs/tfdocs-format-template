BUILD_NAME := tfdocs-format-template

BUILD_DIR     := bin
PLUGIN_FOLDER ?= ~/.tfdocs.d/plugins

GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

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

.PHONY: all clean verify build install help