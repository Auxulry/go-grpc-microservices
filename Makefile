ENGINE=cmd/server/main.go
BUILD_DIR=build

debug:
	go run ${ENGINE} service --svport 3000 --gwport 3001
.PHONY: debug

build:
	@echo "Building app"
	go build -o ${BUILD_DIR}/app ${ENGINE}
	@echo "Success build app. Your app is ready to use in 'build/' directory."
.PHONY: build

proto-gen:
	@echo "Generating the stubs"
	./scripts/proto-gen.sh
	@echo "Success generate stubs. All stubs created are in the 'stubs/' directory"
	@echo "DO NOT EDIT ANY FILES STUBS!"
.PHONY: proto-gen

dependency:
	@echo "Downloading all Go dependencies needed"
	go mod download
	go mod verify
	go mod tidy
	@echo "All Go dependencies was downloaded. you can run 'make debug' to compile locally or 'make build' to build app."
.PHONY: dependency

lint:
	golangci-lint run ./...
.PHONY: lint