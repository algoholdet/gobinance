package binance

// StreamType represent a type of stream at Binance.
type StreamType string

// For your typesafe convenience we provide a number of stream types.
const (
	StreamTypeAggregatedTrade         StreamType = "aggTrade"
	StreamTypeTrade                   StreamType = "trade"
	StreamTypeKLine1m                 StreamType = "kline_1m"
	StreamTypeKLine3m                 StreamType = "kline_3m"
	StreamTypeKLine5m                 StreamType = "kline_5m"
	StreamTypeKLine15m                StreamType = "kline_15m"
	StreamTypeKLine30m                StreamType = "kline_30m"
	StreamTypeKLine1h                 StreamType = "kline_1h"
	StreamTypeKLine2h                 StreamType = "kline_2h"
	StreamTypeKLine4h                 StreamType = "kline_4h"
	StreamTypeKLine6h                 StreamType = "kline_6h"
	StreamTypeKLine8h                 StreamType = "kline_8h"
	StreamTypeKLine12h                StreamType = "kline_12h"
	StreamTypeKLine1d                 StreamType = "kline_1d"
	StreamTypeKLine3d                 StreamType = "kline_3d"
	StreamTypeKLine1w                 StreamType = "kline_1w"
	StreamTypeKLine1M                 StreamType = "kline_1M"
	StreamTypeTicker                  StreamType = "ticker"
	StreamTypeAllMarkedsMarketTickers StreamType = "arr"
	StreamTypePartialDepth5           StreamType = "depth5"
	StreamTypePartialDepth10          StreamType = "depth10"
	StreamTypePartialDepth20          StreamType = "depth20"
	StreamTypeDepth                   StreamType = "depth"
)

func (t StreamType) iface() interface{} {
	switch t {
	case StreamTypeAggregatedTrade:
		return new(AggregatedTrades)

	case StreamTypeTrade:
		return new(Trade)

	case StreamTypeKLine1m, StreamTypeKLine3m, StreamTypeKLine5m,
		StreamTypeKLine15m, StreamTypeKLine30m, StreamTypeKLine1h,
		StreamTypeKLine2h, StreamTypeKLine4h, StreamTypeKLine6h,
		StreamTypeKLine8h, StreamTypeKLine12h, StreamTypeKLine1d,
		StreamTypeKLine3d, StreamTypeKLine1w, StreamTypeKLine1M:

	case StreamTypeTicker:

	case StreamTypeAllMarkedsMarketTickers:

	case StreamTypePartialDepth5, StreamTypePartialDepth10, StreamTypePartialDepth20:

	case StreamTypeDepth:
	}

	return nil
}
