BINARY_NAME=pomocli
BUILD_DIR=./build
INSTALL_DIR=/usr/local/bin

.PHONY: build run install clean

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

run: build
	$(BUILD_DIR)/$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

install: build
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)

# Catch-all to ignore extra arguments
%:
	@:
