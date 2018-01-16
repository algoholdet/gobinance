package binance

import (
	"strings"
)

type Symbol string

// upperCase will return s in uppercase. For some reason some endpoints
// require the symbol in uppercase.
func (s Symbol) upperCase() string {
	return strings.ToUpper(string(s))
}
