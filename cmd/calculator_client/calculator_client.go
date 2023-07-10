package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/radu2020/lovoo/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")

	operandA = flag.Int("a", 1, "operand a")
	operandB = flag.Int("b", 1, "operand b")

	errOperandANotFound  = `operand "a" not found`
	errOperandBNotFound  = `operand "b" not found`
	errOperatorNotFound  = `operator not found`
	errOperatorIncorrect = `invalid operator. Use one of "add", "subtract", "multiply" or "divide"`
)

// isFlagPassed checks if a flag is present
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// IsValidOperator validates the operator.
func IsValidOperator(op string) bool {
	switch op {
	case "add", "subtract", "multiply", "divide":
		return true
	}
	return false
}

type Operation struct {
	OperandA int32
	OperandB int32
	Operator string
}

// FlagsInit parses the CLI flags and argument, checks for missing flags,
// validates the input and returns an Operation or an error.
func FlagsInit() (Operation, error) {
	// Parse CLI flags
	flag.Parse()

	if !isFlagPassed("a") {
		return Operation{}, fmt.Errorf(errOperandANotFound)
	}

	if !isFlagPassed("b") {
		return Operation{}, fmt.Errorf(errOperandBNotFound)
	}

	if len(flag.Args()) < 1 {
		return Operation{}, fmt.Errorf(errOperatorNotFound)
	}

	operator := flag.Args()[0]
	if !IsValidOperator(operator) {
		return Operation{}, fmt.Errorf(errOperatorIncorrect)
	}

	return Operation{
		OperandA: int32(*operandA),
		OperandB: int32(*operandB),
		Operator: operator,
	}, nil
}

// executeCommand creates a gRPC connection to the server, sends the received
// CLI input to be computed and prints out the response to Stdout.
func executeCommand(op Operation) error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Compute(ctx, &pb.ComputeRequest{
		OperandA: op.OperandA,
		OperandB: op.OperandB,
		Operator: op.Operator,
	})
	if err != nil {
		log.Fatalf("Could not compute: %v", err)
	}
	fmt.Println(r.GetResult())
	return nil
}

// Package main implements a client for Calculator service.
func main() {
	op, err := FlagsInit()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Contact the server and print out its response.
	err = executeCommand(op)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
