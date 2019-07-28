package binance

import (
	"encoding/json"
	"fmt"

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

// AggregatedTradesStream will open a websocket stream that will stream
// aggregated trades for symbol. You can use the Read() method when reading
// from the stream. You should call Close() when done.
func (c *Client) AggregatedTradesStream(symbol Symbol) (*AggregatedTradesStream, error) {
	URL := fmt.Sprintf("%s/ws/%s@aggTrade", c.streamBaseURL, string(symbol.LowerCase()))

	conn, err := websocket.Dial(URL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	stream := &AggregatedTradesStream{
		Conn: conn,
	}

	return stream, nil
}
