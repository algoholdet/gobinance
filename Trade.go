package binance

// Trade is a trade matched by the engine.
type Trade struct {
	Symbol Symbol `json:"s"`

	// These match what's seen on a HistoricalTrade.
	TradeID    int64 `json:"t"`
	Price      Value `json:"p"`
	Quantity   Value `json:"q"`
	Timestamp  Time  `json:"T"`
	BuyerMaker bool  `json:"m"`
	BestMatch  bool  `json:"M"`

	BuyerOrderID  int64 `json:"b"`
	SellerOrderID int64 `json:"a"`
}
