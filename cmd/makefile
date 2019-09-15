# Go parameters
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=rpi

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ../
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)