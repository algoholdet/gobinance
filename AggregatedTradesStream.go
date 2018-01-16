package binance

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

// AggregatedTradesStream represents a stream from the aggregated trades endpoint.
type AggregatedTradesStream struct {
	*websocket.Conn
}

// Read a trade from the stream. This will block until a trade is ready.
func (s *AggregatedTradesStream) Read() (*AggregatedTrades, error) {
	var msg = make([]byte, 512)

	n, err := s.Conn.Read(msg)
	if err != nil {
		return nil, err
	}

	trade := &AggregatedTrades{}
	err = json.Unmarshal(msg[:n], trade)
	if err != nil {
		return nil, err
	}

	return trade, nil
}
