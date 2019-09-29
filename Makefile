# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=service
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build_scan clean
build_scan:
	$(GOBUILD) -o $(BINARY_NAME)_scan -v scan_service/main.go
build_notification:
	$(GOBUILD) -o $(BINARY_NAME)_notification -v notification_service/main.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)_notification
	rm -f $(BINARY_NAME)_scan
	rm -f $(BINARY_UNIX)_notification
	rm -f $(BINARY_UNIX)_scan