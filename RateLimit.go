package binance

// RateLimit describes a rate limit at the exchange.
type RateLimit struct {
	Type     string `json:"rateLimitType"` // FIXME: type
	Interval string `json:"interval"`      // FIXME: Type
	Limit    int64  `json:"limit"`
}
