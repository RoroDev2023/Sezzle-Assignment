package calculator

import (
	"errors"
	"math"
)

type Operation string

const (
	Add        Operation = "add"
	Subtract   Operation = "subtract"
	Multiply   Operation = "multiply"
	Divide     Operation = "divide"
	Power      Operation = "power"
	SquareRoot Operation = "sqrt"
	Percentage Operation = "percentage"
)

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrDivisionByZero   = errors.New("division by zero")
	ErrNegativeSqrt     = errors.New("square root requires a non-negative operand")
	ErrInvalidNumber    = errors.New("result is not a finite number")
)

type Calculation struct {
	Operation Operation `json:"operation"`
	A         float64   `json:"a"`
	B         *float64  `json:"b,omitempty"`
}

func Operations() []Operation {
	return []Operation{Add, Subtract, Multiply, Divide, Power, SquareRoot, Percentage}
}

func Calculate(input Calculation) (float64, error) {
	var result float64

	switch input.Operation {
	case Add:
		result = input.A + value(input.B)
	case Subtract:
		result = input.A - value(input.B)
	case Multiply:
		result = input.A * value(input.B)
	case Divide:
		divisor := value(input.B)
		if divisor == 0 {
			return 0, ErrDivisionByZero
		}
		result = input.A / divisor
	case Power:
		result = math.Pow(input.A, value(input.B))
	case SquareRoot:
		if input.A < 0 {
			return 0, ErrNegativeSqrt
		}
		result = math.Sqrt(input.A)
	case Percentage:
		result = input.A * value(input.B) / 100
	default:
		return 0, ErrInvalidOperation
	}

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, ErrInvalidNumber
	}

	return result, nil
}

func RequiresSecondOperand(operation Operation) bool {
	return operation != SquareRoot
}

func value(input *float64) float64 {
	if input == nil {
		return 0
	}
	return *input
}

