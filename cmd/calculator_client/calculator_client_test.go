package main

import (
	"fmt"
	"os"
	"testing"
)

type flagsInput struct {
	operandA         string
	operandB         string
	operator         string
	operandANotFound bool
	operandBNotFound bool
	operatorNotFound bool
}

func TestMainFunc(t *testing.T) {
	type expectation struct {
		out Operation
		err error
	}

	tests := map[string]struct {
		in       flagsInput
		expected expectation
	}{
		"Must_Succeed_Add": {
			in: flagsInput{
				operandA: "6",
				operandB: "2",
				operator: "add",
			},
			expected: expectation{
				out: Operation{
					OperandA: 6,
					OperandB: 2,
					Operator: "add",
				},
				err: nil,
			},
		},
		"Must_Succeed_Subtract": {
			in: flagsInput{
				operandA: "6",
				operandB: "2",
				operator: "subtract",
			},
			expected: expectation{
				out: Operation{
					OperandA: 6,
					OperandB: 2,
					Operator: "subtract",
				},
				err: nil,
			},
		},
		"Must_Succeed_Multiply": {
			in: flagsInput{
				operandA: "6",
				operandB: "2",
				operator: "multiply",
			},
			expected: expectation{
				out: Operation{
					OperandA: 6,
					OperandB: 2,
					Operator: "multiply",
				},
				err: nil,
			},
		},
		"Must_Succeed_Divide": {
			in: flagsInput{
				operandA: "6",
				operandB: "2",
				operator: "divide",
			},
			expected: expectation{
				out: Operation{
					OperandA: 6,
					OperandB: 2,
					Operator: "divide",
				},
				err: nil,
			},
		},
		"Must_Fail_Operand_A_Not_Found": {
			in: flagsInput{
				operandB:         "2",
				operator:         "add",
				operandANotFound: true,
			},
			expected: expectation{
				out: Operation{},
				err: fmt.Errorf(errOperandANotFound),
			},
		},
		"Must_Fail_Operand_B_Not_Found": {
			in: flagsInput{
				operandA:         "6",
				operator:         "add",
				operandBNotFound: true,
			},
			expected: expectation{
				out: Operation{},
				err: fmt.Errorf(errOperandBNotFound),
			},
		},
		"Must_Fail_Operator_Not_Found": {
			in: flagsInput{
				operandA:         "6",
				operandB:         "2",
				operatorNotFound: true,
			},
			expected: expectation{
				out: Operation{},
				err: fmt.Errorf(errOperatorNotFound),
			},
		},
		"Must_Fail_Operator_Incorrect": {
			in: flagsInput{
				operandA: "6",
				operandB: "2",
				operator: "addition",
			},
			expected: expectation{
				out: Operation{},
				err: fmt.Errorf(errOperatorIncorrect),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			// When initial cli flags are missing, skip the output tests
			runOutputTests := true

			// Set Args
			if tt.in.operandANotFound {
				runOutputTests = false
				os.Args = []string{"calculator_client.go", "-b", tt.in.operandB, tt.in.operator}
			} else if tt.in.operandBNotFound {
				runOutputTests = false
				os.Args = []string{"calculator_client.go", "-a", tt.in.operandA, tt.in.operator}
			} else if tt.in.operatorNotFound {
				runOutputTests = false
				os.Args = []string{"calculator_client.go", "-a", tt.in.operandA, "-b", tt.in.operandB}
			} else {
				runOutputTests = false
				os.Args = []string{"calculator_client.go", "-a", tt.in.operandA, "-b", tt.in.operandB, tt.in.operator}
			}

			// Initialize flags
			op, err := FlagsInit()

			// Check for errors
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else if runOutputTests {
				if tt.expected.out.Operator != op.Operator {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.Operator, op.Operator)
				}
				if tt.expected.out.OperandA != op.OperandA {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.OperandA, op.OperandA)
				}
				if tt.expected.out.OperandB != op.OperandB {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.OperandB, op.OperandB)
				}
			}
		})
	}
}
