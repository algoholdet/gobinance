package binance

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

// TradeStream represents a stream from the trades endpoint.
type TradeStream struct {
	*websocket.Conn
}

// Read a trade from the stream. This will block until a trade is ready.
func (s *TradeStream) Read() (*Trade, error) {
	var msg = make([]byte, 512)

	n, err := s.Conn.Read(msg)
	if err != nil {
		return nil, err
	}

	trade := &Trade{}
	err = json.Unmarshal(msg[:n], trade)
	if err != nil {
		return nil, err
	}

	return trade, nil
}
