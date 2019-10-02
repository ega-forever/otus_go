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
build_docker:
	docker build -f notification_Dockerfile -t egorzuev/notification_service:1.0.0 .
	docker build -f scan_Dockerfile -t egorzuev/scan_service:1.0.0 .
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)_notification
	rm -f $(BINARY_NAME)_scan
