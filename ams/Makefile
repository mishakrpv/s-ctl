APP_NAME = ams
SRC_DIR = ./cmd/$(APP_NAME)
BUILD_DIR = ./bin
GO_FILES = $(shell find . -name '*.go')

.PHONY: build
build: $(GO_FILES)
	@mkdir -p $(BUILD_DIR)

	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

.PHONY: run
run: build
	@$(BUILD_DIR)/$(APP_NAME)