CLIENT_BINARY_NAME=calculator_client.out
SERVER_BINARY_NAME=calculator_server.out

# all compiles, runs all the tests and removes the binaries
.PHONY: all
all: build test clean

.PHONY: build
build:
	go build -o ${CLIENT_BINARY_NAME} ./cmd/calculator_client/calculator_client.go
	go build -o ${SERVER_BINARY_NAME} ./cmd/calculator_server/calculator_server.go

# test runs both server and client unit tests
.PHONY: test
test: test-server test-client

.PHONY: test-server
test-server:
	go test -v ./cmd/calculator_server/

.PHONY: test-client
test-client:
	go test -v ./cmd/calculator_client/

.PHONY: clean
clean:
	go clean
	go clean -testcache
	rm ${CLIENT_BINARY_NAME}
	rm ${SERVER_BINARY_NAME}

# run-example starts the grpc server, sends an example client request and outputs the result.
# How to run: `$ make -j2 run-example`
.PHONY: run-example
run-example: run-server run-client

.PHONY: run-server
run-server:
	@go run cmd/calculator_server/calculator_server.go &
	@sleep 3

.PHONY: run-client
run-client:
	@sleep 5
	@go run cmd/calculator_client/calculator_client.go -a 20 -b 4 divide

# run-api generates the api files
.PHONY: run-api
run-api:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        api/calculator.proto
