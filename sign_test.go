package binance

import (
	"testing"
)

// Test sign function bases on example in API documentation.
func TestSignString(t *testing.T) {
	in := "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"

	key := "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
	expected := "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"

	out := signString(in, key)

	if out != expected {
		t.Errorf("Got %s, expected %s", out, expected)
	}
}
