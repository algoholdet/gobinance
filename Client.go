package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

// Client represents a client talking to the Binance API.
type Client struct {
	apiKey        string
	apiSecret     string
	streamBaseURL string
	baseURL       string
	client        *http.Client
}

func APIKey(apiKey string) func(*Client) {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func APISecret(apiSecret string) func(*Client) {
	return func(c *Client) {
		c.apiSecret = apiSecret
	}
}

// BaseURL will define a new base URL. You would probably never use this.
func BaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// StreamBaseURL changes the base URL for the streaming endpoints. You would
// probably never use this.
func StreamBaseURL(streamBaseURL string) func(*Client) {
	return func(c *Client) {
		c.streamBaseURL = streamBaseURL
	}
}

// HTTPClient will change the HTTP client to use. Default is
// http.DefaultClient. This can be used for example if you would like to use
// the AppEngine http client.
func HTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.client = client
	}
}

// NewClient will return a client usable for accessing the Binance API.
func NewClient(options ...func(*Client)) (*Client, error) {
	client := &Client{
		baseURL:       "https://api.binance.com",
		streamBaseURL: "wss://stream.binance.com:9443",
		client:        http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	return client, nil
}

func param(key string, value interface{}) func(url.Values) {
	return func(v url.Values) {
		switch value.(type) {
		case string:
			v.Add(key, value.(string))
		case int:
			v.Add(key, strconv.Itoa(value.(int)))
		default:
			panic("unsupported value type")
		}
	}
}

func (c *Client) publicGet(target interface{}, URI string, params ...func(url.Values)) error {
	URL, _ := url.Parse(fmt.Sprintf("%s%s", c.baseURL, URI))

	values := url.Values{}
	for _, p := range params {
		p(values)
	}

	response, err := c.client.Get(URL.String())
	if err != nil {
		// We should probably do some error parsing here. At least determine
		// if it's a retriable error.
		return err
	}
	defer response.Body.Close()

	if target == nil {
		return nil
	}

	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(target)
}

// Ping will ping the Binance API and return a RTT duration and an error if
// something went wrong.
func (c *Client) Ping() (time.Duration, error) {
	t := time.Now()
	err := c.publicGet(nil, "/api/v1/ping")
	duration := time.Since(t)
	if err != nil {
		return duration, err
	}

	return duration, nil
}

// ServerTime will return the time according to Binance.
func (c *Client) ServerTime() (Time, error) {
	var proxy struct {
		Time Time `json:"serverTime"`
	}

	err := c.publicGet(&proxy, "/api/v1/time")

	return proxy.Time, err
}

// ExchangeInfo returns current exchange trading rules and symbol information.
func (c *Client) ExchangeInfo() (*ExchangeInfo, error) {
	info := &ExchangeInfo{}

	err := c.publicGet(info, "/api/v1/exchangeInfo")

	return info, err
}

func (c *Client) AggregatedTradesStream(symbol Symbol) (*AggregatedTradesStream, error) {
	URL := fmt.Sprintf("%s/ws/%s@aggTrade", c.streamBaseURL, string(symbol))

	conn, err := websocket.Dial(URL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	stream := &AggregatedTradesStream{
		Conn: conn,
	}

	return stream, nil
}

func (c *Client) CombinedStream(streams []StreamID) (*CombinedStream, error) {
	URL := fmt.Sprintf("%s/stream?streams=%s", c.streamBaseURL, joinStreamID(streams))

	conn, err := websocket.Dial(URL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	stream := &CombinedStream{
		Conn: conn,
	}

	return stream, nil
}
