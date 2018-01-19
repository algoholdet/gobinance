package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// signString will sign a string using a method suitable for request signing.
func signString(in string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))

	mac.Write([]byte(in))

	return fmt.Sprintf("%x", (mac.Sum(nil)))
}
