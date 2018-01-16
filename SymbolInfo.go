package binance

// SymbolInfo describes various details about a trading symbol.
type SymbolInfo struct {
	Symbol              Symbol      `json:"symbol"`
	Status              string      `json:"status"`    // FIXME: type
	BaseAsset           string      `json:"baseAsset"` // FIXME: type
	BaseAssetPrecision  int         `json:"baseAssetPrecision"`
	QuoteAsset          string      `json:"quoteAsset"` // FIXME: type
	QuoteAssetPrecision int         `json:"quotePrecision"`
	OrderTypes          []OrderType `json:"orderTypes"`
	AllowIceberg        bool        `json:"icebergAllowed"`
	//Filters             []Filter    `json:"filters"` // FIXME: Implement
}
