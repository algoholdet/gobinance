package binance

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

// CombinedStream is a stream emitting diferent event types.
type CombinedStream struct {
	*websocket.Conn
}

// Read a trade from the stream. This will block until a trade is ready.
func (s *CombinedStream) Read() (interface{}, error) {
	var data = make([]byte, 1024)

	n, err := s.Conn.Read(data)
	if err != nil {
		return nil, err
	}

	return combinedEvent(data[:n])
}

func combinedEvent(data []byte) (interface{}, error) {
	type proxy struct {
		Stream StreamID        `json:"stream"`
		Data   json.RawMessage `json:"data"`
	}

	p := &proxy{}
	err := json.Unmarshal(data, p)
	if err != nil {
		return nil, err
	}

	target := p.Stream.Type().iface()
	if target == nil {
		return nil, fmt.Errorf("Unknown stream type: %s", p.Stream.Type())
	}

	err = json.Unmarshal(p.Data, target)
	if err != nil {
		return nil, err
	}

	return target, nil
}
