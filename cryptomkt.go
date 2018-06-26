package cryptomkt

import (
	"net/http"
)

// Client client that hold Services.
type Client struct {
	PaymentService
	PublicService
	PrivateService
}

// NewClient instance a new cruptomkt client.
func NewClient(APIKey, secret string) *Client {
	priClient := &httpClient{
		client: &http.Client{},
		key:    APIKey,
		secret: secret,
	}

	pubClient := &httpClient{
		client: &http.Client{},
		key:    APIKey,
		secret: secret,
	}

	return &Client{
		PaymentService: PaymentService{priClient, true},
		PublicService:  PublicService{pubClient, false},
		PrivateService: PrivateService{priClient, true},
	}
}

// NewPublicClient expose only public endpoints.
func NewPublicClient() *Client {
	pubClient := &httpClient{
		client: &http.Client{},
	}

	return &Client{
		PublicService: PublicService{pubClient, false},
	}
}
