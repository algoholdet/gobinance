package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// signString will sign a string using a method suitable for request signing.
func signString(in string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))

	// This should never fail. Famous last words.
	_, _ = mac.Write([]byte(in))

	return fmt.Sprintf("%x", (mac.Sum(nil)))
}
