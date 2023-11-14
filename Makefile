run: build
	@./bin/message-queue

build: 
	@go build -o bin/message-queue

test: 
	@go test -v ./...
