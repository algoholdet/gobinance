package binance

import (
	"strings"
)

// Symbol it a trading pair. For example "BTCUSDT". This is not case-sensitive.
type Symbol string

// upperCase will return s in uppercase. For some reason some endpoints
// require the symbol in uppercase.
func (s Symbol) upperCase() string {
	return strings.ToUpper(string(s))
}

func (s Symbol) lowerCase() string {
	return strings.ToLower(string(s))
}
