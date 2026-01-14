BINARY_NAME=mctextgen
SOURCE=cmd/mc-gen/main.go
LDFLAGS=-ldflags="-s -w"
build:
	go build $(LDFLAGS) -o $(BINARY_NAME) $(SOURCE)
run:
	go run $(SOURCE)
clean:
	go clean
	rm -f $(BINARY_NAME)*
deps:
	go mod download