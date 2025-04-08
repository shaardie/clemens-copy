APP_NAME := clemens-copy
SRC := main.go
BIN_DIR := bin

.PHONY: all clean build windows linux mac

all: clean build

build: windows linux mac

windows:
	@echo "🔧 Baue Windows-Binary..."
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME).exe $(SRC)

linux:
	@echo "🐧 Baue Linux-Binary..."
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)_linux $(SRC)

mac:
	@echo "🍏 Baue macOS-Binary..."
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)_mac $(SRC)

clean:
	@echo "🧹 Entferne alte Builds..."
	rm -rf $(BIN_DIR)
	mkdir -p $(BIN_DIR)
