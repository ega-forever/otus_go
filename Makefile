# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build clean
build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
compile_proto:
	protoc --go_out=plugins=grpc:internal/protocol/grpc api/*.proto
run:
	$(GOBUILD) -o $(BINARY_NAME) -v main.go
	./$(BINARY_NAME)