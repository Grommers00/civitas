.PHONY: build

GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main

# TODO: Add tests to all once integrated
all: sam

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf build/$(BINARY_NAME)

run: build
	./build/$(BINARY_NAME)

deps:
	$(GOGET) -v all 

sam:
	sam build
