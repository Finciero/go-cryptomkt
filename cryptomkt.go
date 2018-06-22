package cryptomkt

import (
	"net/http"
)

type Client struct {
	PaymentService
}

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
