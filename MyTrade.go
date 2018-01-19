package binance

// MyTrade is a trade order in the Binance system.
type MyTrade struct {
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Price           Value  `json:"price"`
	Quantity        Value  `json:"qty"`
	Commission      Value  `json:"commission"`
	CommissionAsset string `json:"commissionAsset"` // FIXME: type
	TimeStamp       Time   `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}
