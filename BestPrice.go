package binance

// BestPrice is used to describe the best price/quantity in the order book.
type BestPrice struct {
	Bid OrderBookPoint `json:"bid"`
	Ask OrderBookPoint `json:"ask"`
}

type bestPriceProxy struct {
	Symbol      Symbol `json:"symbol"`
	BidPrice    Value  `json:"bidPrice"`
	BidQuantity Value  `json:"bidQty"`
	AskPrice    Value  `json:"askPrice"`
	AskQuantity Value  `json:"askQty"`
}

// real will convert p to an BestPrice.
func (p *bestPriceProxy) real() (*BestPrice, error) {
	bestPrice := &BestPrice{
		Bid: OrderBookPoint{
			Price:    p.BidPrice,
			Quantity: p.BidQuantity,
		},
		Ask: OrderBookPoint{
			Price:    p.AskPrice,
			Quantity: p.AskQuantity,
		},
	}

	return bestPrice, nil
}
