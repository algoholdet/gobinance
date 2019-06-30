package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
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
	dumpWriter    io.Writer
}

// APIKey will parse the API key to the client. This is not needed for all
// calls.
func APIKey(apiKey string) func(*Client) {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// APISecret will parse the API secret to the client. This is not needed for
// all calls.
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

// DumpWriter will instruct Client to dump all HTTP requests and responses to
// and from Binance to w.
func DumpWriter(w io.Writer) func(*Client) {
	return func(c *Client) {
		c.dumpWriter = w
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
		case int64:
			v.Add(key, strconv.FormatInt(value.(int64), 10))
		default:
			panic(fmt.Sprintf("unsupported value type: %T", value))
		}
	}
}

func (c *Client) buildRequest(method string, URI string, params ...func(url.Values)) (*http.Request, error) {
	values := url.Values{}
	for _, p := range params {
		p(values)
	}

	URL := fmt.Sprintf("%s%s?%s", c.baseURL, URI, values.Encode())

	return http.NewRequest(method, URL, nil)
}

func (c *Client) doRequest(target interface{}, req *http.Request) error {
	if c.dumpWriter != nil {
		r, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.dumpWriter, "HTTP Request:\n%s\n", string(r))
	}

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if c.dumpWriter != nil {
		r, err := httputil.DumpResponse(response, true)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.dumpWriter, "HTTP Response:\n%s\n", string(r))
	}

	// FIXME: Handle various errors from the API here

	if target == nil {
		return nil
	}

	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(target)
}

func (c *Client) publicGet(target interface{}, URI string, params ...func(url.Values)) error {
	req, _ := c.buildRequest("GET", URI, params...)

	return c.doRequest(target, req)
}

func (c *Client) marketGet(target interface{}, URI string, params ...func(url.Values)) error {
	if c.apiKey == "" {
		return errors.New("no API key set")
	}

	req, _ := c.buildRequest("GET", URI, params...)

	req.Header.Add("X-MBX-APIKEY", c.apiKey)

	return c.doRequest(target, req)
}

func (c *Client) signedCall(target interface{}, method string, URI string, params ...func(url.Values)) error {
	if c.apiSecret == "" {
		return errors.New("no API secret set")
	}

	// Add a timestamp to the request.
	timestamp := fmt.Sprintf("%d",
		time.Now().UnixNano()/int64(time.Millisecond))

	params = append(params, param("timestamp", timestamp))

	req, _ := c.buildRequest("GET", URI, params...)

	// Add a signature to the request. It will be safe to simply add it here
	// using '&', since timestamp will always be set and we will never
	// encounter an empty query string.
	signature := signString(req.URL.RawQuery, c.apiSecret)
	req.URL.RawQuery += "&signature=" + signature

	req.Header.Add("X-MBX-APIKEY", c.apiKey)

	return c.doRequest(target, req)
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

// OrderBook will return the current order book for symbol.
func (c *Client) OrderBook(symbol Symbol, limit int) (*OrderBook, error) {
	proxy := &orderBookProxy{}

	err := c.publicGet(proxy, "/api/v1/depth",
		param("symbol", symbol.UpperCase()),
		param("limit", limit),
	)
	if err != nil {
		return nil, err
	}

	return proxy.real()
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

// HistoricalTrades retrieves historical trades for symbol. You can use
// Limit() and FromID().
func (c *Client) HistoricalTrades(symbol Symbol, options ...QueryFunc) ([]HistoricalTrade, error) {
	var trades []HistoricalTrade

	err := c.marketGet(&trades, "/api/v1/historicalTrades",
		param("symbol", symbol.UpperCase()),
		newQuery(options).params(),
	)

	return trades, err
}

// LatestPriceAll will retrieve latest price for all symbols.
func (c *Client) LatestPriceAll() (map[Symbol]Value, error) {
	var proxy []struct {
		Symbol Symbol `json:"symbol"`
		Price  Value  `json:"price"`
	}

	err := c.publicGet(&proxy, "/api/v3/ticker/price")
	if err != nil {
		return nil, err
	}

	m := make(map[Symbol]Value, len(proxy))

	for _, p := range proxy {
		m[p.Symbol] = p.Price
	}

	return m, nil
}

// LatestPrice will retrieve the latest price for a symbol.
func (c *Client) LatestPrice(symbol Symbol) (Value, error) {
	var proxy struct {
		Symbol Symbol `json:"symbol"`
		Price  Value  `json:"price"`
	}

	err := c.publicGet(&proxy, "/api/v3/ticker/price",
		param("symbol", symbol.UpperCase()),
	)
	if err != nil {
		return "-1.0", err
	}

	return proxy.Price, nil
}

// BestPriceAll returns the best price/quantity for all symbols.
func (c *Client) BestPriceAll() (map[Symbol]BestPrice, error) {
	var proxy []bestPriceProxy

	err := c.publicGet(&proxy, "/api/v3/ticker/bookTicker")
	if err != nil {
		return nil, err
	}

	bestPrices := make(map[Symbol]BestPrice, len(proxy))

	for _, p := range proxy {
		bestPrice, _ := p.real()

		bestPrices[p.Symbol] = *bestPrice
	}

	return bestPrices, nil
}

// BestPrice returns best price/qty on the order book for a symbol.
func (c *Client) BestPrice(symbol Symbol) (*BestPrice, error) {
	var proxy bestPriceProxy
	err := c.publicGet(&proxy, "/api/v3/ticker/bookTicker",
		param("symbol", symbol.UpperCase()),
	)
	if err != nil {
		return nil, err
	}

	return proxy.real()
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

// ChangeStatisticsAll returns 24 hour price change statistics for all symbols.
func (c *Client) ChangeStatisticsAll() (map[Symbol]ChangeStatistics, error) {
	var proxy []ChangeStatistics

	err := c.publicGet(&proxy, "/api/v1/ticker/24hr")
	if err != nil {
		return nil, err
	}

	stats := make(map[Symbol]ChangeStatistics, len(proxy))

	for _, p := range proxy {
		stats[p.Symbol] = p
	}

	return stats, nil
}

// ChangeStatistics returns 24 hour price change statistics for symbol.
func (c *Client) ChangeStatistics(symbol Symbol) (*ChangeStatistics, error) {
	var changeStatistics ChangeStatistics

	err := c.publicGet(&changeStatistics, "/api/v1/ticker/24hr",
		param("symbol", symbol.UpperCase()),
	)
	if err != nil {
		return nil, err
	}

	return &changeStatistics, nil
}

// MyTrades return trades for a specific symbol. You can refine the query with
// Limit() & FromID().
// Note: recvWindow parameter not supported (yet).
func (c *Client) MyTrades(symbol Symbol, options ...QueryFunc) ([]MyTrade, error) {
	var orders []MyTrade

	err := c.signedCall(&orders, "GET", "/api/v3/myTrades",
		param("symbol", symbol.UpperCase()),
		newQuery(options).params(),
	)
	if err != nil {
		return nil, err
	}

	return orders, nil
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

// TradeStream will open a websocket stream that will stream trades for symbol.
// You can use the Read() method when reading from the stream. You should call
// Close() when done.
func (c *Client) TradeStream(symbol Symbol) (*TradeStream, error) {
	URL := fmt.Sprintf("%s/ws/%s@trade", c.streamBaseURL, string(symbol.LowerCase()))

	conn, err := websocket.Dial(URL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	stream := &TradeStream{
		Conn: conn,
	}

	return stream, nil
}

// CombinedStream will open a websocket stream for multiple events.
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
