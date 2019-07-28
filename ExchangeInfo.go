package binance

// ExchangeInfo describes various details about the exchange configuration.
type ExchangeInfo struct {
	Timezone   string       `json:"timezone"`
	ServerTime Time         `json:"serverTime"`
	RateLimits []RateLimit  `json:"rateLimits"`
	Symbols    []SymbolInfo `json:"symbols"`
	//ExchangeFilters []json.RawMessage `json:"exchangeFilters"` // FIXME: What is this?
}

// ExchangeInfo returns current exchange trading rules and symbol information.
func (c *Client) ExchangeInfo() (*ExchangeInfo, error) {
	info := &ExchangeInfo{}

	err := c.publicGet(info, "/api/v1/exchangeInfo")

	return info, err
}
