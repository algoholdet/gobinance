package binance

import (
	"strings"
)

// StreamID identifies a stream from Binance.
type StreamID string

// Type returns the type of a stream.
func (s StreamID) Type() StreamType {
	parts := strings.Split(string(s), "@")
	if len(parts) != 2 {
		return ""
	}

	return StreamType(parts[1])
}

// Symbol returns the symbol of a stream.
func (s StreamID) Symbol() Symbol {
	parts := strings.Split(string(s), "@")
	if len(parts) != 2 {
		return ""
	}

	return Symbol(parts[0])
}

// NewStreamID will return a new StreamID consisting of a symbol and a stream
// type.
func NewStreamID(symbol Symbol, typ StreamType) StreamID {
	return StreamID(string(symbol.lowerCase()) + "@" + string(typ))
}

// joinStreamID is a helper to join a slice of StreamID's suited for passing
// to the combined streams API.
// Lifted from strings.Join.
func joinStreamID(a []StreamID) StreamID {
	const sep = StreamID("/")

	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return StreamID(b)
}
