# lovoo

Client-Server application, which can calculate simple math tasks.

## Installation instructions

### Clone the repository
```shell
$ git clone https://github.com/radu2020/lovoo.git
```

### Install Go 1.20.1
For installation instructions, see Go’s [Getting Started](https://go.dev/doc/install) guide.

### Install Protocol buffer compiler, protoc, version 3.
For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).
 
### Install Go plugins for the protocol compiler
Install the protocol compiler plugins for Go using the following commands:

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Update your PATH so that the protoc compiler can find the plugins:
```shell
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

### How to generate the proto files
>All `make` commands should be run from the project root directory.

```shell
$ make generate-api
```

## Usage
1. Start the gRPC Server.
2. Send a gRPC Client request to the server.
Here we need to specify operands `a`, `b` and `operator` as part of the CLI flags (examples below).
3. Output is printed to Client Stdout.

### How to start the server 
```shell
$ go run cmd/calculator_server/calculator_server.go
```

### How to send the client request (example)
Here we divide 20 by 4.
```shell
$ go run cmd/calculator_client/calculator_client.go -a 20 -b 4 divide
$ 5
```

### Supported commands
1. `add` performs addition (a+b)
2. `subtract` performs subtraction (a-b)
3. `multiply` performs multiplication (a*b)
4. `divide` performs division (a/b)

### How to run the e2e example
For this to work the server should not be already running (port :50051 should be free).
This starts the server, sends a predefined client request, and outputs the result.
Use `Ctrl + C` to terminate the process.
```shell
$ make -j2 run-example
$ 5
```

### Build binaries, run unit tests, and clean up
```shell
make all
```

### Run Server unit tests
The server unit tests test the add, subtract, multiply and divide computations.
```shell
make test-server
```

### Run Client unit tests
The client unit tests test input CLI flags such as the operands and the operator.
```shell
make test-client
```

### Run all unit tests
Run both client and server unit tests using command:
```shell
make test
```

Tests Output:
```shell
go test -v ./cmd/calculator_server/
=== RUN   TestCalculatorServer_Compute
=== RUN   TestCalculatorServer_Compute/Must_Succeed_Add
=== RUN   TestCalculatorServer_Compute/Must_Succeed_Subtract
=== RUN   TestCalculatorServer_Compute/Must_Succeed_Multiply
=== RUN   TestCalculatorServer_Compute/Must_Succeed_Division
--- PASS: TestCalculatorServer_Compute (0.00s)
    --- PASS: TestCalculatorServer_Compute/Must_Succeed_Add (0.00s)
    --- PASS: TestCalculatorServer_Compute/Must_Succeed_Subtract (0.00s)
    --- PASS: TestCalculatorServer_Compute/Must_Succeed_Multiply (0.00s)
    --- PASS: TestCalculatorServer_Compute/Must_Succeed_Division (0.00s)
PASS
ok      github.com/radu2020/lovoo/cmd/calculator_server 0.126s
go test -v ./cmd/calculator_client/
=== RUN   TestMainFunc
=== RUN   TestMainFunc/Must_Fail_Operand_B_Not_Found
=== RUN   TestMainFunc/Must_Fail_Operator_Not_Found
=== RUN   TestMainFunc/Must_Fail_Operator_Incorrect
=== RUN   TestMainFunc/Must_Succeed_Add
=== RUN   TestMainFunc/Must_Succeed_Subtract
=== RUN   TestMainFunc/Must_Succeed_Multiply
=== RUN   TestMainFunc/Must_Succeed_Divide
=== RUN   TestMainFunc/Must_Fail_Operand_A_Not_Found
--- PASS: TestMainFunc (0.00s)
    --- PASS: TestMainFunc/Must_Fail_Operand_B_Not_Found (0.00s)
    --- PASS: TestMainFunc/Must_Fail_Operator_Not_Found (0.00s)
    --- PASS: TestMainFunc/Must_Fail_Operator_Incorrect (0.00s)
    --- PASS: TestMainFunc/Must_Succeed_Add (0.00s)
    --- PASS: TestMainFunc/Must_Succeed_Subtract (0.00s)
    --- PASS: TestMainFunc/Must_Succeed_Multiply (0.00s)
    --- PASS: TestMainFunc/Must_Succeed_Divide (0.00s)
    --- PASS: TestMainFunc/Must_Fail_Operand_A_Not_Found (0.00s)
PASS
ok      github.com/radu2020/lovoo/cmd/calculator_client 0.125s
```

## Thank you!
Looking forward to see you soon!