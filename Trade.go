package binance

// Trade represents a trade in the past.
type Trade struct {
	ID         int64 `json:"id"`
	Price      Value `json:"price"`
	Quantity   Value `json:"qty"`
	Timestamp  Time  `json:"time"`
	BuyerMaker bool  `json:"isBuyerMaker"`
	BestMatch  bool  `json:"isBestMatch"`
}
