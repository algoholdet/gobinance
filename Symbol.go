package binance

import (
	"strings"
)

// Symbol it a trading pair. For example "BTCUSDT". This is not case-sensitive.
type Symbol string

const zeroSymbol Symbol = ""

// UpperCase will return s in uppercase. For some reason some endpoints
// require the symbol in uppercase.
func (s Symbol) UpperCase() string {
	return strings.ToUpper(string(s))
}

// LowerCase will return s in lowercase. For some reason some endpoints
// require the symbol in lowercase.
func (s Symbol) LowerCase() string {
	return strings.ToLower(string(s))
}
