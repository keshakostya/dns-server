BINARY  := dns-server
PKG		:= ./cmd/$(BINARY)
BIN_DIR := bin


build:
	go build -o bin/dns-server $(PKG)