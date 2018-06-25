package cryptomkt

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// PublicService represent the implementation of Cryptomkt's service for public endpoints.
type PublicService struct {
	client  *httpClient
	Private bool
}

// MarketResponse represent the market response.
type MarketResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

// GetMarkets returns a list of available markets
func (ps *PublicService) GetMarkets() (*MarketResponse, error) {
	resp, err := ps.client.get("/market", nil)
	if err != nil {
		return nil, err
	}

	var rr MarketResponse
	if err := unmarshalJSON(resp.Body, &rr); err != nil {
		return nil, err
	}

	return &rr, nil
}

// Ticker represent a ticker.
type Ticker struct {
	High      string `json:"high"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	LastPrice string `json:"last_price"`
	Volume    string `json:"volume"`
	Market    string `json:"market"`
	Timestamp string `json:"time[[[[stamp"`
}

// TickerResponse represent a ticker response.
type TickerResponse struct {
	Status string    `json:"status,omitempty"`
	Data   []*Ticker `json:"data,omitempty"`
}

// GetTicker returns a list of available ticker
func (ps *PublicService) GetTicker(market string) (*TickerResponse, error) {
	resp, err := ps.client.get(fmt.Sprintf("/ticker?market=%s", market), nil)
	if err != nil {
		return nil, err
	}

	var tr TickerResponse
	if err := unmarshalJSON(resp.Body, &tr); err != nil {
		return nil, err
	}

	return &tr, nil
}

// Book represent a Exchange market order.
type Book struct {
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

// Pagination represent a pagination info.
type Pagination struct {
	Previous int `json:"previous,omitempty"`
	Limit    int `json:"limit,omitempty"`
	Page     int `json:"page,omitempty"`
	Next     int `json:"next,omitempty"`
}

// BooksResponse represent a ticker.
type BooksResponse struct {
	Status     string      `json:"status"`
	Data       []*Book     `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// BooksOptions represent query params for book request.
type BooksOptions struct {
	Market string `json:"market,omitempty" url:"market"`
	Kind   string `json:"kind,omitempty" url:"kind"`
	Page   int    `json:"page,omitempty" url:"page,omitempty"`
	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`
}

// GetOrdersBook return a collection of active orders.
func (ps *PublicService) GetOrdersBook(opts *BooksOptions) (*BooksResponse, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := ps.client.get(fmt.Sprintf("/book?%s", v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var br BooksResponse
	if err := unmarshalJSON(resp.Body, &br); err != nil {
		return nil, err
	}

	return &br, nil
}

// Trade represent a trade.
type Trade struct {
	// Tipo de transacción. buy o sell
	MarketTaker string `json:"market_taker,omitempty"`
	// Precio al cual se realizó la transacción
	Price string `json:"price,omitempty"`
	// Cantidad de la transacción
	Amount string `json:"amount,omitempty"`
	// ID de la transacción
	Tid string `json:"tid,omitempty"`
	// Fecha de la transacción
	Timestamp string `json:"timestamp,omitempty"`
	// Par de mercado donde se realizó la transacción
	Market string `json:"market,omitempty"`
}

// TradesResponse represrtns a trades response.
type TradesResponse struct {
	Status     string      `json:"status"`
	Data       []*Trade    `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// TradesOptions represent query params for trades request.
type TradesOptions struct {
	Market    string `json:"market,omitempty" url:"market"`
	StartDate string `json:"start,omitempty" url:"start"`
	EndDate   string `json:"end,omitempty" url:"end"`
	Page      int    `json:"page,omitempty" url:"page,omitempty"`
	Limit     int    `json:"limit,omitempty" url:"limit,omitempty"`
}

// GetTrades return a collection of trades.
func (ps *PublicService) GetTrades(opts *TradesOptions) (*TradesResponse, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := ps.client.get(fmt.Sprintf("/trades?%s", v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var tr TradesResponse
	if err := unmarshalJSON(resp.Body, &tr); err != nil {
		return nil, err
	}

	return &tr, nil
}
