package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/radu2020/lovoo/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.CalculatorClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterCalculatorServer(baseServer, NewServer())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewCalculatorClient(conn)

	return client, closer
}

func TestCalculatorServer_Compute(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.ComputeResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.ComputeRequest
		expected expectation
	}{
		"Must_Succeed_Add": {
			in: &pb.ComputeRequest{
				OperandA: 6,
				OperandB: 2,
				Operator: "add",
			},
			expected: expectation{
				out: &pb.ComputeResponse{
					Result: 8,
				},
				err: nil,
			},
		},
		"Must_Succeed_Subtract": {
			in: &pb.ComputeRequest{
				OperandA: 6,
				OperandB: 2,
				Operator: "subtract",
			},
			expected: expectation{
				out: &pb.ComputeResponse{
					Result: 4,
				},
				err: nil,
			},
		},
		"Must_Succeed_Multiply": {
			in: &pb.ComputeRequest{
				OperandA: 6,
				OperandB: 2,
				Operator: "multiply",
			},
			expected: expectation{
				out: &pb.ComputeResponse{
					Result: 12,
				},
				err: nil,
			},
		},
		"Must_Succeed_Division": {
			in: &pb.ComputeRequest{
				OperandA: 6,
				OperandB: 2,
				Operator: "divide",
			},
			expected: expectation{
				out: &pb.ComputeResponse{
					Result: 3,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Compute(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Result != out.Result {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.Result, out.Result)
				}
			}
		})
	}
}
