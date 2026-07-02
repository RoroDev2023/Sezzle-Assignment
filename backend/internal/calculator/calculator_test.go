package calculator

import "testing"

func TestCalculate(t *testing.T) {
	t.Parallel()

	two := 2.0
	three := 3.0
	ten := 10.0
	twenty := 20.0

	tests := []struct {
		name      string
		input     Calculation
		want      float64
		wantError error
	}{
		{name: "add", input: Calculation{Operation: Add, A: 2, B: &three}, want: 5},
		{name: "subtract", input: Calculation{Operation: Subtract, A: 5, B: &three}, want: 2},
		{name: "multiply", input: Calculation{Operation: Multiply, A: 5, B: &three}, want: 15},
		{name: "divide", input: Calculation{Operation: Divide, A: 10, B: &two}, want: 5},
		{name: "power", input: Calculation{Operation: Power, A: 2, B: &three}, want: 8},
		{name: "sqrt", input: Calculation{Operation: SquareRoot, A: 9}, want: 3},
		{name: "percentage", input: Calculation{Operation: Percentage, A: 20, B: &ten}, want: 2},
		{name: "division by zero", input: Calculation{Operation: Divide, A: 10, B: pointer(0)}, wantError: ErrDivisionByZero},
		{name: "negative sqrt", input: Calculation{Operation: SquareRoot, A: -1}, wantError: ErrNegativeSqrt},
		{name: "unknown operation", input: Calculation{Operation: "mod", A: 10, B: &twenty}, wantError: ErrInvalidOperation},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Calculate(tt.input)
			if tt.wantError != nil {
				if err != tt.wantError {
					t.Fatalf("Calculate() error = %v, want %v", err, tt.wantError)
				}
				return
			}

			if err != nil {
				t.Fatalf("Calculate() unexpected error = %v", err)
			}

			if got != tt.want {
				t.Fatalf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func pointer(value float64) *float64 {
	return &value
}

