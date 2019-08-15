package binance

import (
	"strconv"
)

// Value represents a value from Binance. Value is stored as a string to avoid
// rounding errors.
type Value string

const zeroValue Value = ""

// Float64 returns the value as a float64.
func (v Value) Float64() float64 {
	f, _ := strconv.ParseFloat(string(v), 64)
	return f
}

// Float64Err will return two values. The value as a float and an error that
// will be nil if the value could be decoded as a float64.
func (v Value) Float64Err() (float64, error) {
	return strconv.ParseFloat(string(v), 64)
}

// String implements Stringer.
func (v Value) String() string {
	return string(v)
}
