run: build
	@./bin/message-queue

build: 
	@go build -o bin/message-queue

test: 
	@go test -v ./...

testp: 
	@go run cmd/testpublish/main.go

testc:
	@go run cmd/testconsumer/main.go
