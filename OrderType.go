package binance

import (
	"encoding/json"
	"fmt"
)

type OrderType string

const (
	Limit  OrderType = "LIMIT"
	Market OrderType = "MARKET"
)

// UnmarshalJSON implements json.Unmarshaler while making sure only enums
// that we know about end up in a OrderType variable.
func (o *OrderType) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	order := OrderType(s)

	switch order {
	case Limit, Market:
		*o = order
	default:
		return fmt.Errorf("%s is not a valid order type", s)
	}

	return nil
}

// String implement Stringer.
func (o OrderType) String() string {
	return string(o)
}
