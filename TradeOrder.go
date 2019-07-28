package binance

// TradeOrder is a trade order in the Binance system.
type TradeOrder struct {
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Price           Value  `json:"price"`
	Quantity        Value  `json:"qty"`
	Commission      Value  `json:"commission"`
	CommissionAsset string `json:"commissionAsset"` // FIXME: type
	TimeStamp       Time   `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

// MyTrades return trades for a specific symbol. You can refine the query with
// Limit() & FromID().
// Note: recvWindow parameter not supported (yet).
func (c *Client) MyTrades(symbol Symbol, options ...QueryFunc) ([]TradeOrder, error) {
	var orders []TradeOrder

	err := c.signedCall(&orders, "GET", "/api/v3/myTrades",
		param("symbol", symbol.UpperCase()),
		newQuery(options).params(),
	)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
