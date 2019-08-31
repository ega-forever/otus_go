GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=telnet_util
ROOT_PATH=./cmd
BINARY_UNIX=$(BINARY_NAME)_unix
offset=0
limit=0

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) $(ROOT_PATH)
clean:
	rm -f $(BINARY_NAME)