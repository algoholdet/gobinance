package binance

import (
	"encoding/json"
)

// CandleStick is a single "candle" in a candle stick graph.
type CandleStick struct {
	OpenTime                 Time  `json:"openTime"`
	Open                     Value `json:"open"`
	High                     Value `json:"high"`
	Low                      Value `json:"low"`
	Close                    Value `json:"close"`
	Volume                   Value `json:"volume"`
	CloseTime                Time  `json:"closeTime"`
	QuoteAssetVolume         Value `json:"quoteAssetVolume"`
	NumberOfTrades           int   `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  Value `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume Value `json:"takerBuyBaseAssetVolume"`
}

// candleStickProxy is used for unmarshalling JSON from the API. We have to use
// this because Binance gives a an array (!) with all values.
type candleStickProxy [12]json.RawMessage

// real will return a CandleStick based on p.
func (p candleStickProxy) real() (*CandleStick, error) {
	candleStick := CandleStick{}

	targets := []interface{}{
		&candleStick.OpenTime,
		&candleStick.Open,
		&candleStick.High,
		&candleStick.Low,
		&candleStick.Close,
		&candleStick.Volume,
		&candleStick.CloseTime,
		&candleStick.QuoteAssetVolume,
		&candleStick.NumberOfTrades,
		&candleStick.TakerBuyBaseAssetVolume,
		&candleStick.TakerBuyQuoteAssetVolume,
	}

	for i, target := range targets {
		err := json.Unmarshal(p[i], target)
		if err != nil {
			return nil, err
		}
	}

	return &candleStick, nil
}
