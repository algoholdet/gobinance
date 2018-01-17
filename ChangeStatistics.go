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
