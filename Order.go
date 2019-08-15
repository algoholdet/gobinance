package binance

import (
	"errors"
	"net/url"
)

// Order describes an order in the Binance Exchange.
type Order struct {
	Symbol                   Symbol      `json:"symbol"`
	ID                       int         `json:"orderId"`
	OrderListID              int         `json:"orderListId"`
	ClientOrderID            string      `json:"clientOrderId"`
	Price                    Value       `json:"price"`
	Quantity                 Value       `json:"origQty"`
	ExecutedQuantity         Value       `json:"executedQty"`
	CummulativeQuoteQuantity Value       `json:"cummulativeQuoteQty"`
	Status                   OrderStatus `json:"status"`
	TimeInForce              TimeInForce `json:"timeInForce"`
	Type                     OrderType   `json:"type"`
	Side                     OrderSide   `json:"side"`
	StopPrice                Value       `json:"stopPrice"`
	IcebergQuantity          Value       `json:"icebergQty"`
	Time                     Time        `json:"time"`
	Updated                  Time        `json:"updateTime"`
	Working                  bool        `json:"isWorking"`
}

// SubmitOrder will submit order for processing.
func (c *Client) SubmitOrder(order *Order) error {
	return c.submitOrder("/api/v3/order", order)
}

// SubmitTestOrder will submit a test order.
func (c *Client) SubmitTestOrder(order *Order) error {
	return c.submitOrder("/api/v3/order/test", order)
}

func (c *Client) submitOrder(uri string, order *Order) error {
	var result interface{}

	params := []func(url.Values){
		param("newOrderRespType", "RESULT"),
		param("symbol", order.Symbol),
		param("side", order.Side),
		param("type", order.Type),
		param("quantity", order.Quantity),
	}

	if order.TimeInForce != zeroTimeInForceZero {
		params = append(params, param("timeInForce", order.TimeInForce))
	}

	if order.Price != zeroValue {
		params = append(params, param("price", order.Price))
	}

	if order.StopPrice != zeroValue {
		params = append(params, param("stopPrice", order.StopPrice))
	}

	if order.ID != 0 {
		params = append(params, param("newClientOrderId", order.ID))
	}

	if order.IcebergQuantity != zeroValue {
		params = append(params, param("icebergQty", order.IcebergQuantity))
	}

	err := c.signedCall(&result, "POST", uri, params...)
	if err != nil {
		return err
	}

	return nil
}

// CancelOrder cancels a live order.
func (c *Client) CancelOrder(symbol Symbol, clientOrderID string, id int) (*Order, error) {
	params := []func(url.Values){
		param("symbol", symbol),
	}

	if clientOrderID == "" && id == 0 {
		return nil, errors.New("clientOrderID and ID empty")
	}

	if clientOrderID != "" && id != 0 {
		return nil, errors.New("clientOrderID and ID both set")
	}

	if clientOrderID != "" {
		params = append(params, param("origClientOrderId", clientOrderID))
	}

	if id != 0 {
		params = append(params, param("orderId", id))
	}

	var order Order
	err := c.signedCall(&order, "DELETE", "/api/v3/order", params...)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// OrderStatus queries the status of an order.
func (c *Client) OrderStatus(symbol Symbol, clientOrderID string, id int) (*Order, error) {
	params := []func(url.Values){
		param("symbol", symbol),
	}

	if clientOrderID == "" && id == 0 {
		return nil, errors.New("clientOrderID and ID empty")
	}

	if clientOrderID != "" && id != 0 {
		return nil, errors.New("clientOrderID and ID both set")
	}

	if clientOrderID != "" {
		params = append(params, param("origClientOrderId", clientOrderID))
	}

	if id != 0 {
		params = append(params, param("orderId", id))
	}

	var order Order
	err := c.signedCall(&order, "GET", "/api/v3/order", params...)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// OpenOrders lists the currently open orders.
func (c *Client) OpenOrders(symbol Symbol) ([]Order, error) {
	params := []func(url.Values){}

	if symbol != zeroSymbol {
		params = append(params, param("symbol", symbol))
	}

	results := make([]Order, 0, 100)
	err := c.signedCall(&results, "GET", "/api/v3/openOrders", params...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// AllOrders will list all orders open or closed.
func (c *Client) AllOrders(symbol Symbol) ([]Order, error) {
	results := make([]Order, 0, 100)
	err := c.signedCall(&results, "GET", "/api/v3/allOrders", param("symbol", symbol))
	if err != nil {
		return nil, err
	}

	return results, nil
}
