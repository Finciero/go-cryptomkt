package cryptomkt

import (
	"net/http"
	"net/url"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.cryptomkt.com",
	Path:   "v1",
}

type Client struct {
	PaymentService
}

func NewClient(APIKey string) *Client {
	hclient := httpClient{
		client: &http.Client{},
		secret: APIKey,
	}

	return &Client{
		PaymentService: PaymentService{&hclient},
	}
}
