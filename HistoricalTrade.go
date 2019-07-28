package binance

// HistoricalTrade represents a trade in the past.
type HistoricalTrade struct {
	TradeID       int64 `json:"id"`
	Price         Value `json:"price"`
	Quantity      Value `json:"qty"`
	QuoteQuantity Value `json:"quoteQty"`
	Timestamp     Time  `json:"time"`
	BuyerMaker    bool  `json:"isBuyerMaker"`
	BestMatch     bool  `json:"isBestMatch"`
}

// HistoricalTrades retrieves historical trades for symbol. You can use
// Limit() and FromID().
func (c *Client) HistoricalTrades(symbol Symbol, options ...QueryFunc) ([]HistoricalTrade, error) {
	var trades []HistoricalTrade

	err := c.marketGet(&trades, "/api/v1/historicalTrades",
		param("symbol", symbol.UpperCase()),
		newQuery(options).params(),
	)

	return trades, err
}
