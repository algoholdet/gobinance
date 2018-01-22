package binance

import (
	"errors"
)

// OrderBook represents the current order book for a specific symbol.
type OrderBook struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []OrderBookPoint
	Asks         []OrderBookPoint
}

// orderBookProxy is a type used as a proxy for the format from the API. The
// format is not easible convertable to Go structs. This type helps somewhat.
type orderBookProxy struct {
	LastUpdateID int64           `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

// OrderBookPoint is a small helper for OrderBook.
type OrderBookPoint struct {
	Price    Value
	Quantity Value
}

// real will convert p to an OrderBook.
func (p *orderBookProxy) real() (*OrderBook, error) {
	convert := func(in [][]interface{}, out []OrderBookPoint) error {
		for i, b := range in {
			if len(b) == 3 {
				p, ok := b[0].(string)
				if !ok {
					return errors.New("Unknown format")
				}

				q, ok := b[1].(string)
				if !ok {
					return errors.New("Unknown format")
				}

				out[i].Price = Value(p)
				out[i].Quantity = Value(q)
			}
		}

		return nil
	}

	orderbook := &OrderBook{
		LastUpdateID: p.LastUpdateID,
		Bids:         make([]OrderBookPoint, len(p.Bids), len(p.Bids)),
		Asks:         make([]OrderBookPoint, len(p.Asks), len(p.Asks)),
	}

	err := convert(p.Bids, orderbook.Bids)
	if err != nil {
		return nil, err
	}

	err = convert(p.Asks, orderbook.Asks)
	if err != nil {
		return nil, err
	}

	return orderbook, nil
}
