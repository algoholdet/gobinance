package binance

// ExchangeInfo describes various details about the exchange configuration.
type ExchangeInfo struct {
	Timezone   string       `json:"timezone"`
	ServerTime Time         `json:"serverTime"`
	RateLimits []RateLimit  `json:"rateLimits"`
	Symbols    []SymbolInfo `json:"symbols"`
	//ExchangeFilters []json.RawMessage `json:"exchangeFilters"` // FIXME: What is this?
}
