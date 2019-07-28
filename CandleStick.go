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

// CandleStick returns Kline/candlestick bars for symbol. Klines are uniquely
// identified by their open time. You can refine the query with Limit(),
// StartTime() and EndTime().
func (c *Client) CandleStick(symbol Symbol, interval string, options ...QueryFunc) ([]CandleStick, error) {
	var proxy []candleStickProxy

	err := c.publicGet(&proxy, "/api/v1/klines",
		param("symbol", symbol.UpperCase()),
		param("interval", interval),
		newQuery(options).params(),
	)
	if err != nil {
		return nil, err
	}

	sticks := make([]CandleStick, len(proxy), len(proxy))
	for i, p := range proxy {
		stick, err := p.real()
		if err != nil {
			return nil, err
		}

		sticks[i] = *stick
	}

	return sticks, nil
}
