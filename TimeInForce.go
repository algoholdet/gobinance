package binance

import (
	"encoding/json"
	"fmt"
)

// TimeInForce is describing for how long a trade should remain active
// or executed.
type TimeInForce string

// The different states of an order.
const (
	GoodTillCancelled TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"
	FillOrKill        TimeInForce = "FOK"

	GTC = GoodTillCancelled
	IOC = ImmediateOrCancel
	FOK = FillOrKill

	zeroTimeInForceZero = ""
)

// UnmarshalJSON implements json.Unmarshaler while making sure only enums
// that we know about end up in a TimeInForce variable.
func (t *TimeInForce) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	status := TimeInForce(s)

	switch status {
	case GTC, IOC, FOK:
		*t = status
	default:
		return fmt.Errorf("%s is not a valid TimeInForce", s)
	}

	return nil
}

// String implement Stringer.
func (t TimeInForce) String() string {
	return string(t)
}
