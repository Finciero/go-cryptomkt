package cryptomkt

import (
	"net/http"
)

// Client client that hold Services.
type Client struct {
	PaymentService
}

// NewClient instance a new cruptomkt client.
func NewClient(APIKey, secret string) *Client {
	hclient := httpClient{
		client: &http.Client{},
		key:    APIKey,
		secret: secret,
	}

	return &Client{
		PaymentService: PaymentService{&hclient},
	}
}
