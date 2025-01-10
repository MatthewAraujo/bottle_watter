BIN_DIR=bin

LINUX_BINARY=$(BIN_DIR)/watter_bottle-linux-amd64
DARWIN_BINARY=$(BIN_DIR)/watter_bottle-darwin-amd64
WINDOWS_BINARY=$(BIN_DIR)/watter_bottle.exe

all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY)

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(DARWIN_BINARY)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BINARY)

clean:
	rm -rf $(BIN_DIR)