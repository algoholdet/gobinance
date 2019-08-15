package binance

// AccountInfo describes the current account.
type AccountInfo struct {
	MakerCommission  int  `json:"makerCommission"`
	TakerCommission  int  `json:"takerCommission"`
	BuyerCommission  int  `json:"buyerCommission"`
	SellerCommission int  `json:"sellerCommission"`
	CanTrade         bool `json:"canTrade"`
	CanWithdraw      bool `json:"canWithdraw"`
	CanDeposit       bool `json:"canDeposit"`
	Balances         []struct {
		Asset  string `json:"asset"`
		Free   Value  `json:"free"`
		Locked Value  `json:"locked"`
	} `json:"balances"`
}

// AccountInfo retrieves various information about the account.
func (c *Client) AccountInfo() (*AccountInfo, error) {
	var info AccountInfo
	err := c.signedCall(&info, "GET", "/api/v3/account")
	if err != nil {
		return nil, err
	}

	return &info, nil
}
