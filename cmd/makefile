# Go parameters
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=rpi

all: test build

build:
	env GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ../
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
publish:
	scp rpi pi@192.168.178.23:/home/pi/
