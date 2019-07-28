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

// BestPriceAll returns the best price/quantity for all symbols.
func (c *Client) BestPriceAll() (map[Symbol]BestPrice, error) {
	var proxy []bestPriceProxy

	err := c.publicGet(&proxy, "/api/v3/ticker/bookTicker")
	if err != nil {
		return nil, err
	}

	bestPrices := make(map[Symbol]BestPrice, len(proxy))

	for _, p := range proxy {
		bestPrice, _ := p.real()

		bestPrices[p.Symbol] = *bestPrice
	}

	return bestPrices, nil
}

// BestPrice returns best price/qty on the order book for a symbol.
func (c *Client) BestPrice(symbol Symbol) (*BestPrice, error) {
	var proxy bestPriceProxy
	err := c.publicGet(&proxy, "/api/v3/ticker/bookTicker",
		param("symbol", symbol.UpperCase()),
	)
	if err != nil {
		return nil, err
	}

	return proxy.real()
}
