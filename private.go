package cryptomkt

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// PrivateService represent the implementation of Cryptomkt's service for private endpoints.
type PrivateService struct {
	client  *httpClient
	Private bool
}

// OrderAmount represent an Amount in MakerOrder
type OrderAmount struct {
	// Cantidad original de la orden
	Original float64 `json:"original,string,omitempty"`
	// Cantidad restante de la orden. Solo en órdenes activas
	Remaining float64 `json:"remaining,string,omitempty"`
	// Cantidad ejecutada de la orden. Solo en órdenes ejecutadas
	Executed float64 `json:"executed,string,omitempty"`
}

// MarketOrder represent market order response.
type MarketOrder struct {
	// ID de la orden
	ID string `json:"id,omitempty"`
	// Estado de la orden. active o executed
	Status string `json:"status,omitempty"`
	// Tipo de orden. buy o sell
	Type string `json:"type,omitempty"`
	// Precio límite de la orden
	Price int64 `json:"price,string,omitempty"`
	//
	Amount *OrderAmount `json:"amount,omitempty"`
	//  Precio de ejecución
	ExecutionPrice float64 `json:"execution_price,string,omitempty"`
	// Precio de ejecución promedio ponderado. 0 si no se ejecuta.
	AvgExecutionPrice int64 `json:"avg_execution_price,string,omitempty"`
	// Par de mercado
	Market string `json:"market,omitempty"`
	// Fecha de creación
	CreatedAt string `json:"created_at,omitempty"`
	// Fecha de actualización. Solo en órdenes activas
	UpdatedAt string `json:"updated_at,omitempty"`
	// Fecha de ejecución. Solo en órdenes ejecutadas
	ExecutedAt string `json:"executed_at,omitempty"`
}

// MarketOrdersResponse represents a collection of market orders.
type MarketOrdersResponse struct {
	Status     string         `json:"status,omitempty"`
	Data       []*MarketOrder `json:"data,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty"`
}

// MarketOrderOptions represents marker order query options
type MarketOrderOptions struct {
	// Par de mercado
	Market string `url:"market"`
	// Página a consultar
	Page int `url:"page,omitempty"`
	// Límite de objetos por página. Por defecto es 20. Mínimo 20 , máximo 100
	Limit int `url:"limit,omitempty"`
}

// GetActiveOrders return a collection of active orders.
func (ps *PrivateService) GetActiveOrders(opts *MarketOrderOptions) (*MarketOrdersResponse, error) {
	ps.client.SetPrivate(ps.Private)
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := ps.client.get(fmt.Sprintf("/orders/active?%s", v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var mor MarketOrdersResponse
	if err := unmarshalJSON(resp.Body, &mor); err != nil {
		return nil, err
	}

	return &mor, nil
}

// GetExecutedOrders return a collection of active orders.
func (ps *PrivateService) GetExecutedOrders(opts *MarketOrderOptions) (*MarketOrdersResponse, error) {
	ps.client.SetPrivate(ps.Private)
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := ps.client.get(fmt.Sprintf("/orders/executed?%s", v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var mor MarketOrdersResponse
	if err := unmarshalJSON(resp.Body, &mor); err != nil {
		return nil, err
	}

	return &mor, nil
}

// MarketOrderRequest represents a market order request.
type MarketOrderRequest struct {
	Market string  `json:"market,omitempty"`
	Amount float64 `json:"amount,string,omitempty"`
	Price  int     `json:"price,string,omitempty"`
	Type   string  `json:"type,omitempty"`
}

// Params returns a map used to sign the requests
func (mor *MarketOrderRequest) Params() url.Values {
	form := url.Values{}

	form.Add("amount", fmt.Sprintf("%f", mor.Amount))
	form.Add("market", mor.Market)
	form.Add("price", fmt.Sprintf("%d", mor.Price))
	form.Add("type", mor.Type)

	return form
}

// MarketOrderResponse represents a market order response
type MarketOrderResponse struct {
	Status string       `json:"status,omitempty"`
	Data   *MarketOrder `json:"data,omitempty"`
}

// CreateOrder creates a new order.
func (ps *PrivateService) CreateOrder(mor *MarketOrderRequest) (*MarketOrderResponse, error) {
	ps.client.SetPrivate(ps.Private)
	resp, err := ps.client.postForm("/orders", mor.Params())
	if err != nil {
		return nil, err
	}

	var morr MarketOrderResponse
	if err := unmarshalJSON(resp.Body, &morr); err != nil {
		return nil, err
	}

	return &morr, nil
}

// OrderStatusOption represents an order status request.
type OrderStatusOption struct {
	ID string `url:"id"`
}

// GetOrderStatus return an market order
func (ps *PrivateService) GetOrderStatus(opts *OrderStatusOption) (*MarketOrderResponse, error) {
	ps.client.SetPrivate(ps.Private)
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := ps.client.get(fmt.Sprintf("/orders/status?%s", v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	var morr MarketOrderResponse
	if err := unmarshalJSON(resp.Body, &morr); err != nil {
		return nil, err
	}
	return &morr, nil
}

// CancelOrderRequest represents an order cancel request.
type CancelOrderRequest struct {
	ID string `json:"id,omitempty"`
}

// Params returns a map used to sign the requests.
func (cor *CancelOrderRequest) Params() url.Values {
	form := url.Values{}

	form.Add("id", cor.ID)

	return form
}

// CancelOrder cancel an order.
func (ps *PrivateService) CancelOrder(mor *CancelOrderRequest) (*MarketOrderResponse, error) {
	ps.client.SetPrivate(ps.Private)
	resp, err := ps.client.postForm("/orders/cancel", mor.Params())
	if err != nil {
		return nil, err
	}

	var morr MarketOrderResponse
	if err := unmarshalJSON(resp.Body, &morr); err != nil {
		return nil, err
	}

	return &morr, nil
}

// Balance represents a wallet balance.
type Balance struct {
	Wallet    string  `json:"wallet,omitempty"`
	Available float64 `json:"available,string,omitempty"`
	Balance   float64 `json:"balance,string,omitempty"`
}

// BalanceResponse represents a balance response.
type BalanceResponse struct {
	Status string     `json:"status,omitempty"`
	Data   []*Balance `json:"data,omitempty"`
}

// GetBalance returns balance from wallets of the given key.
func (ps *PrivateService) GetBalance() (*BalanceResponse, error) {
	ps.client.SetPrivate(ps.Private)
	resp, err := ps.client.get("/balance", nil)
	if err != nil {
		return nil, err
	}

	var br BalanceResponse
	if err := unmarshalJSON(resp.Body, &br); err != nil {
		return nil, err
	}

	return &br, nil
}
