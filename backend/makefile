# Go Params

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main

# TODO: Add tests to all once integrated
all: build

build:
	$(GOBUILD) -o build/$(BINARY_NAME) -v cmd/main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf build/$(BINARY_NAME)

run: build
	./build/$(BINARY_NAME)

deps:
	$(GOGET) -v all