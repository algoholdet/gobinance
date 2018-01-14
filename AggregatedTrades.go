package binance

import (
	"fmt"
)

// AggregatedTrades is a trade at the Binance exchance.
type AggregatedTrades struct {
	TradeID      int64   `json:"a"`
	Price        float64 `json:"p,string"`
	Quantity     float64 `json:"q,string"`
	FirstTradeID int64   `json:"f"`
	LastTradeID  int64   `json:"l"`
	Timestamp    Time    `json:"T"`
	Maker        bool    `json:"m"`
	BestMatch    bool    `json:"M"`
}

// ColorString will format the trade as a string suitable for printing to the
// console. If you provide previous it will try to color the price red or
// green.
func (t *AggregatedTrades) ColorString(previous *AggregatedTrades) string {
	ret := fmt.Sprintf("%10d ", t.TradeID)
	if previous != nil {
		if previous.Price > t.Price {
			ret += "\033[31m"
		} else if previous.Price < t.Price {
			ret += "\033[32m"
		}
	}

	ret += fmt.Sprintf("%.02f\033[0m %.03f %s", t.Price, t.Quantity, t.Timestamp.String())

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
