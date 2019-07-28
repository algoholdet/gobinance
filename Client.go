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
)

// Client represents a client talking to the Binance API.
type Client struct {
	apiKey        string
	apiSecret     string
	streamBaseURL string
	baseURL       string
	client        *http.Client
	dumpWriter    io.Writer
	usedWeight    int
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

	client.SetOptions(options...)

	return client, nil
}

// SetOptions can be used to set various options on Client.
func (c *Client) SetOptions(options ...func(*Client)) {
	for _, option := range options {
		option(c)
	}
}

func param(key string, value interface{}) func(url.Values) {
	return func(v url.Values) {
		switch t := value.(type) {
		case string:
			v.Add(key, t)
		case int:
			v.Add(key, strconv.Itoa(t))
		case int64:
			v.Add(key, strconv.FormatInt(t, 10))
		default:
			panic(fmt.Sprintf("unsupported value type: %T", value))
		}
	}
}

func (c *Client) buildRequest(method string, uri string, params ...func(url.Values)) (*http.Request, error) {
	values := url.Values{}
	for _, p := range params {
		p(values)
	}

	URL := fmt.Sprintf("%s%s?%s", c.baseURL, uri, values.Encode())

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

	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("got http status code %d", response.StatusCode)
	}

	if uw, err := strconv.Atoi(response.Header.Get("X-Mbx-Used-Weight")); err != nil {
		c.usedWeight = uw
	}

	if target == nil {
		return nil
	}

	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(target)
}

// UsedWeight will return the total weight used in the present minute.
func (c *Client) UsedWeight() int {
	return c.usedWeight
}

func (c *Client) publicGet(target interface{}, uri string, params ...func(url.Values)) error {
	req, _ := c.buildRequest("GET", uri, params...)

	return c.doRequest(target, req)
}

func (c *Client) marketGet(target interface{}, uri string, params ...func(url.Values)) error {
	if c.apiKey == "" {
		return errors.New("no API key set")
	}

	req, _ := c.buildRequest("GET", uri, params...)

	req.Header.Add("X-MBX-APIKEY", c.apiKey)

	return c.doRequest(target, req)
}

func (c *Client) signedCall(target interface{}, method string, uri string, params ...func(url.Values)) error {
	if c.apiSecret == "" {
		return errors.New("no API secret set")
	}

	// Add a timestamp to the request.
	timestamp := fmt.Sprintf("%d",
		time.Now().UnixNano()/int64(time.Millisecond))

	params = append(params, param("timestamp", timestamp))

	req, _ := c.buildRequest(method, uri, params...)

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
