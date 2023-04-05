# You can change file env like .env or .env.*.local
include .env

ENGINE=cmd/server/main.go
BUILD_DIR=build
CONN_STRING="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}"
MIGRATION_PATH=db/migrations
RPC_PORT=4000
GATEWAY_PORT=4001

debug:
	go run ${ENGINE} service --svport ${RPC_PORT} --gwport ${GATEWAY_PORT}
.PHONY: debug

build:
	@echo "Building app"
	go build -o ${BUILD_DIR}/app ${ENGINE}
	@echo "Success build app. Your app is ready to use in 'build/' directory."
.PHONY: build

proto-gen:	clean-proto
	@echo "Generating the stubs"
	./scripts/proto-gen.sh
	@echo "Success generate stubs. All stubs created are in the 'stubs/' directory"
	@echo "Generating the Swagger UI"
	./scripts/swagger-ui-gen.sh
	@echo "Success generate Swagger UI. If you want to change Swagger UI to previous version copy the previous version from './cache/swagger-ui' directory"
	@echo "You can try swagger-ui with command 'make debug'"
	@echo "DO NOT EDIT ANY FILES STUBS!"
.PHONY: proto-gen

ssl-gen:
	@echo "Generating ssl configuration"
	./scripts/ssl-gen.sh
	@echo "Success generate ssl configuration. All SSL Configuration created in the 'ssl/' directory"
	@echo "DO NOT EXPOSE SSL DIRECTORY!"
.PHONY: ssl-gen

dependency:
	@echo "Downloading all Go dependencies needed"
	go mod download
	go mod verify
	go mod tidy
	@echo "All Go dependencies was downloaded. you can run 'make debug' to compile locally or 'make build' to build app."
.PHONY: dependency

clean-proto:
	@echo "Delete all previous stubs ..."
	rm -rf stubs/*
	@echo "All stubs successfully deleted"
.PHONY: clean-proto

tidy:
	@echo "Synchronize dependency"
	go mod tidy
	@echo "Finish Synchronize dependency"
.PHONY: tidy

lint:
	golangci-lint run ./...
.PHONY: lint

migration-up:
	@echo "Starting migrations up"
	migrate -database ${CONN_STRING} -path ${MIGRATION_PATH} up
	@echo "Finish migration up"
.PHONY: migration-up

migration-down:
	@echo "Starting migrations down"
	migrate -database ${CONN_STRING} -path ${MIGRATION_PATH} down -all
	@echo "Finish migrations down"
.PHONY: migration-down