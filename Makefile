include .env

BUILD_DIR = $(PWD)/dist

.PHONY: dev
dev:
	./bin/air server --port $(APP_PORT)

.PHONY: clean
clean:
	rm -rf ./dist

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

.PHONY: start
start: build
	$(BUILD_DIR)/$(APP_NAME)

.PHONY: update-deps
update-deps:
	go get -u && go mod tidy
