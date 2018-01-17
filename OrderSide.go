package binance

import (
	"encoding/json"
	"fmt"
)

// OrderSide is an enum determined the side of an order.
type OrderSide string

// OrderSide is always OrderSideBuy or OrderSideSell.
const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

// UnmarshalJSON implements json.Unmarshaler while making sure only enums
// that we know about end up in a OrderSide variable.
func (o *OrderSide) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	side := OrderSide(s)

	switch side {
	case OrderSideBuy, OrderSideSell:
		*o = side
	default:
		return fmt.Errorf("%s is not a valid order side", s)
	}

	return nil
}

// String implement Stringer.
func (o OrderSide) String() string {
	return string(o)
}
