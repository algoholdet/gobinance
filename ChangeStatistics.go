package binance

// ChangeStatistics describes a change to a symbol.
type ChangeStatistics struct {
	Symbol                Symbol `jsoin:"symbol"`
	PriceChange           Value  `jsoin:"priceChange"`
	PriceChangePercent    Value  `jsoin:"priceChangePercent"`
	WeightedAveragegPrice Value  `jsoin:"weightedAvgPrice"`
	PreviousClosePrice    Value  `jsoin:"prevClosePrice"`
	LastPrice             Value  `jsoin:"lastPrice"`
	LastQuantity          Value  `jsoin:"lastQty"`
	BidPrice              Value  `jsoin:"bidPrice"`
	AskPrice              Value  `jsoin:"askPrice"`
	PpenPrice             Value  `jsoin:"openPrice"`
	HighPrice             Value  `jsoin:"highPrice"`
	LowPrice              Value  `jsoin:"lowPrice"`
	Volume                Value  `jsoin:"volume"`
	QuoteVolume           Value  `jsoin:"quoteVolume"`
	OpenTime              Time   `jsoin:"openTime"`
	CloseTime             Time   `jsoin:"closeTime"`
	FirstTradeID          int64  `jsoin:"firstId"`
	LastTradeID           int64  `jsoin:"lastId"`
	NumberOfTrades        int    `jsoin:"count"`
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
