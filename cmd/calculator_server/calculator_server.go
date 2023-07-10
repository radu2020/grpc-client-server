// Package main implements a server for Calculator service.
package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/radu2020/lovoo/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func NewServer() pb.CalculatorServer {
	t := &calculatorServer{}
	return t
}

// Compute implements api.Calculator
func (s *calculatorServer) Compute(ctx context.Context, req *pb.ComputeRequest) (*pb.ComputeResponse, error) {
	var result int32

	switch req.Operator {
	case "add":
		result = req.OperandA + req.OperandB
	case "subtract":
		result = req.OperandA - req.OperandB
	case "multiply":
		result = req.OperandA * req.OperandB
	case "divide":
		result = req.OperandA / req.OperandB
	}

	return &pb.ComputeResponse{Result: result}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
