package binance

import (
	"fmt"
)

// AggregatedTrades is a trade at the Binance exchance.
type AggregatedTrades struct {
	Symbol       Symbol `json:"s"`
	TradeID      int64  `json:"a"`
	Price        Value  `json:"p"`
	Quantity     Value  `json:"q"`
	FirstTradeID int64  `json:"f"`
	LastTradeID  int64  `json:"l"`
	Timestamp    Time   `json:"T"`
	Maker        bool   `json:"m"`
	BestMatch    bool   `json:"M"`
}

// ColorString will format the trade as a string suitable for printing to the
// console. If you provide previous it will try to color the price red or
// green.
func (t *AggregatedTrades) ColorString(previous *AggregatedTrades) string {
	ret := fmt.Sprintf("%10s %10d ", t.Symbol, t.TradeID)
	if previous != nil {
		if previous.Price > t.Price {
			ret += "\033[31m"
		} else if previous.Price < t.Price {
			ret += "\033[32m"
		}
	}

	ret += fmt.Sprintf("%.02f\033[0m %.03f %s", t.Price.Float64(), t.Quantity.Float64(), t.Timestamp.String())

	if t.Maker {
		ret += "\033[36m Maker\033[0m"
	} else {
		ret += " Maker"
	}

	if t.BestMatch {
		ret += "\033[36m Bestmatch\033[0m"
	} else {
		ret += " Bestmatch"
	}

	return ret
}

// AggregateTrades will return aggregated historic trades for symbol. You can
// query using FromID(), StartTime(), EndTime() and Limit().
func (c *Client) AggregateTrades(symbol Symbol, options ...QueryFunc) ([]AggregatedTrades, error) {
	var aggTrades []AggregatedTrades

	err := c.publicGet(&aggTrades, "/api/v1/aggTrades",
		param("symbol", symbol.UpperCase()),
		newQuery(options).params(),
	)

	return aggTrades, err
}
