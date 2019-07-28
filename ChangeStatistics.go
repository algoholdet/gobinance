package binance

// ChangeStatistics describes a change to a symbol.
type ChangeStatistics struct {
	Symbol                Symbol `json:"symbol"`
	PriceChange           Value  `json:"priceChange"`
	PriceChangePercent    Value  `json:"priceChangePercent"`
	WeightedAveragegPrice Value  `json:"weightedAvgPrice"`
	PreviousClosePrice    Value  `json:"prevClosePrice"`
	LastPrice             Value  `json:"lastPrice"`
	LastQuantity          Value  `json:"lastQty"`
	BidPrice              Value  `json:"bidPrice"`
	AskPrice              Value  `json:"askPrice"`
	PpenPrice             Value  `json:"openPrice"`
	HighPrice             Value  `json:"highPrice"`
	LowPrice              Value  `json:"lowPrice"`
	Volume                Value  `json:"volume"`
	QuoteVolume           Value  `json:"quoteVolume"`
	OpenTime              Time   `json:"openTime"`
	CloseTime             Time   `json:"closeTime"`
	FirstTradeID          int64  `json:"firstId"`
	LastTradeID           int64  `json:"lastId"`
	NumberOfTrades        int    `json:"count"`
}

// ChangeStatisticsAll returns 24 hour price change statistics for all symbols.
func (c *Client) ChangeStatisticsAll() (map[Symbol]ChangeStatistics, error) {
	var proxy []ChangeStatistics

	err := c.publicGet(&proxy, "/api/v1/ticker/24hr")
	if err != nil {
		return nil, err
	}

	stats := make(map[Symbol]ChangeStatistics, len(proxy))

	for _, p := range proxy {
		stats[p.Symbol] = p
	}

	return stats, nil
}

// ChangeStatistics returns 24 hour price change statistics for symbol.
func (c *Client) ChangeStatistics(symbol Symbol) (*ChangeStatistics, error) {
	var changeStatistics ChangeStatistics

	err := c.publicGet(&changeStatistics, "/api/v1/ticker/24hr",
		param("symbol", symbol.UpperCase()),
	)
	if err != nil {
		return nil, err
	}

	return &changeStatistics, nil
}
