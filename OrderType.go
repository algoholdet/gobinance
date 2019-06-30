package binance

import (
	"encoding/json"
	"fmt"
)

// OrderType is the type of an order.
type OrderType string

// OrderType is a limit or a market order.
const (
	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
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
	case OrderTypeLimit, OrderTypeLimitMaker, OrderTypeMarket, OrderTypeStopLossLimit, OrderTypeTakeProfitLimit:
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
