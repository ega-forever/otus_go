# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=service

all: test build_scan clean
build_scan:
	$(GOBUILD) -o $(BINARY_NAME)_scan -v scan_service/main.go
build_notification:
	$(GOBUILD) -o $(BINARY_NAME)_notification -v notification_service/main.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)_notification
	rm -f $(BINARY_NAME)_scan