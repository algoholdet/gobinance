package binance

import (
	"encoding/json"
	"fmt"
)

type OrderStatus string

const (
	New             OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
	PendingCancel   OrderStatus = "PENDING_CANCEL"
	Rejected        OrderStatus = "REJECTED"
	Expired         OrderStatus = "EXPIRED"
)

// UnmarshalJSON implements json.Unmarshaler while making sure only enums
// that we know about end up in a OrderStatus variable.
func (o *OrderStatus) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	status := OrderStatus(s)

	switch status {
	case New, PartiallyFilled, Filled, Canceled, PendingCancel, Rejected, Expired:
		*o = status
	default:
		return fmt.Errorf("%s is not a valid order status", s)
	}

	return nil
}

// String implement Stringer.
func (o OrderStatus) String() string {
	return string(o)
}
