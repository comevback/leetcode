APP_NAME := leetcode-setup
VERSION := 1.0.0

.PHONY: all clean build

all: clean build

clean:
	rm -f $(APP_NAME)-linux $(APP_NAME).exe $(APP_NAME)-mac

build:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe
	GOOS=darwin GOARCH=arm64 go build -o $(APP_NAME)-mac

run: build
	./$(APP_NAME)
