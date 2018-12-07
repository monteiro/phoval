GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMODGET=$(GOCMD) mod

ENTRYPOINT=cmd/main.go

BINARY_NAME=phoval
BINARY_LINUX=$(BINARY_NAME)-linux
BINARY_MAC=$(BINARY_NAME)-mac
BINARY_WINDOWS=$(BINARY_NAME)-windows.exe

DOCKERCOMPOSE=docker-compose -f docker/docker-compose.yml

#
.PHONY: deps
deps:
	$(GOCMD) mod vendor

# fetch all versions
.PHONY: build
build: build-linux build-osx build-windows

# run the server (when developing)
.PHONY: run
run:
	$(GOCMD) run $(ENTRYPOINT); exit 0;

# run all tests
.PHONY: tests
tests:
	$(GOTEST) -race ./...

# compile linux version
.PHONY: build-linux
build-linux: deps
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -o bin/$(BINARY_LINUX) $(ENTRYPOINT)

# compile osx version
.PHONY: build-osx
build-osx: deps
	CGO_ENABLED=0GOOS=darwin $(GOBUILD) -o bin/$(BINARY_MAC) $(ENTRYPOINT)

# compile windows version
.PHONY: build-windows
build-windows: deps
	CGO_ENABLED=0 GOOS=windows $(GOBUILD) -o bin/$(BINARY_WINDOWS) $(ENTRYPOINT)

# clean all binaries compiled
.PHONY: clean
clean:
	rm -rvf bin/$(BINARY_LINUX)
	rm -rvf bin/$(BINARY_WINDOWS)
	rm -rvf bin/$(BINARY_MAC)

# execute all migrations
.PHONY: migrate
migrate:
	bin/migrate -source file://migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

# create new migration script
# usage: make create-migration MIGRATION_NAME=create_user
.PHONY: create-migration
create-migration:
	bin/migrate create -ext sql -dir migrations $(MIGRATION_NAME)

# build the container inside the Dockerfile
.PHONY:
docker-build:
	$(DOCKERCOMPOSE) build

# run the docker-compose with the database
.PHONY: docker-up
docker-up: docker-build
	$(DOCKERCOMPOSE) up; exit 0;

# run the http server in development mode
.PHONY: docker-run
docker-run: docker-build
	$(DOCKERCOMPOSE) build
