package binance

import (
	"errors"
	"testing"
)

func TestValueFloat64(t *testing.T) {
	expectError := errors.New("error")
	cases := []struct {
		in          string
		expected    float64
		expectedErr error
	}{
		{"10.0", 10.0, nil},
		{"-10.0", -10.0, nil},
		{"10", 10.0, nil},
		{"-10", -10.0, nil},
		{"0.0", 0.0, nil},
		{"a", 0.0, expectError},
		{"", 0.0, expectError},
	}

	for _, c := range cases {
		result, err := Value(c.in).Float64Err()

		if err == nil && c.expectedErr != nil {
			t.Errorf("Float64Err did not return an error")
		}
		if err != nil && c.expectedErr == nil {
			t.Errorf("Float64Err returned an error")
		}

		if result != c.expected {
			t.Errorf("Float64Err(%s) returned %f, expected %f", c.in, result, c.expected)
		}

		if Value(c.in).Float64() != c.expected {
			t.Errorf("Float64(%s) returned %f, expected %f", c.in, result, c.expected)
		}
	}
}
