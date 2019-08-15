# `import "github.com/algoholdet/gobinance"`

This package provides access to the [Binance](https://binance.com/) [API](https://github.com/binance-exchange/binance-official-api-docs).

## To-Do

| Endpoint                          | Security | Status |
|-----------------------------------|----------|--------|
| GET /api/v1/ping                  | Public   | ✓      |
| GET /api/v1/time                  | Public   | ✓      |
| GET /api/v1/exchangeInfo          | Public   | (✓)    |
| GET /api/v1/depth                 | Public   | ✓      |
| GET /api/v1/trades                | Public   |        |
| GET /api/v1/aggTrades             | Public   | ✓      |
| GET /api/v1/historicalTrades      | Key      | ✓      |
| GET /api/v1/klines                | Public   | ✓      |
| GET /api/v3/avgPrice              | Public   |        |
| GET /api/v1/ticker/24hr           | Public   | ✓      |
| GET /api/v3/ticker/price          | Public   | ✓      |
| GET /api/v3/ticker/bookTicker     | Public   | ✓      |
| GET /api/v1/ticker/allPrices      | Public   | ?      |
| GET /api/v1/ticker/allBookTickers | Public   | ?      |
| POST /api/v1/order                | Signed   | ×      |
| POST /api/v3/order                | Signed   | ✓      |
| POST /api/v3/order/test           | Signed   | ✓      |
| GET /api/v3/order                 | Signed   | ✓      |
| DELETE /api/v3/order              | Signed   | ✓      |
| GET /api/v3/openOrders            | Signed   | ✓      |
| GET /api/v3/allOrders             | Signed   | ✓      |
| POST /api/v3/order/oco            | Signed   |        |
| DELETE /api/v3/orderList          | Signed   |        |
| GET /api/v3/orderList             | Signed   |        |
| GET /api/v3/allOrderList          | Signed   |        |
| GET /api/v3/openOrderList         | Signed   |        |
| GET /api/v3/account               | Signed   |        |
| GET /api/v3/myTrades              | Signed   | ✓      |
| POST /wapi/v3/withdraw.html       | Signed   |        |
| GET /wapi/v3/depositHistory.html  | Signed   |        |
| GET /wapi/v3/withdrawHistory.html | Signed   |        |
| GET /wapi/v3/depositAddress.html  | Signed   |        |
| POST /api/v1/userDataStream       | Key      |        |
| PUT /api/v1/userDataStream        | Key      |        |
| DELETE /api/v1/userDataStream     | Key      |        |
| Aggregate Trade Streams           | Public   | ✓      |
| Trade Streams                     | Public   | ✓      |
| Kline/Candlestick Streams         | Public   |        |
| Individual Symbol Ticker Streams  | Public   |        |
| All Market Tickers Stream         | Public   |        |
| Partial Book Depth Streams        | Public   |        |
| Diff. Depth Stream                | Public   |        |
| Combined Stream                   | Public   | (✓)    |
| User Data Websocket               | Key?     |        |
| Error handling                    | All      |        |

(✓): Partially implemented

×: Replaced by newer/REST endpoint.
