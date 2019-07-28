package binance

import (
	"encoding/json"
	"fmt"

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

// TradeStream will open a websocket stream that will stream trades for symbol.
// You can use the Read() method when reading from the stream. You should call
// Close() when done.
func (c *Client) TradeStream(symbol Symbol) (*TradeStream, error) {
	URL := fmt.Sprintf("%s/ws/%s@trade", c.streamBaseURL, symbol.LowerCase())

	conn, err := websocket.Dial(URL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	stream := &TradeStream{
		Conn: conn,
	}

	return stream, nil
}
