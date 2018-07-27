package cryptomkt

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Client client that hold Services.
type Client struct {
	PaymentService
	PublicService
	PrivateService
}

// Debug method to turn on logs.
func (c *Client) Debug() {
	log.SetOutput(os.Stdout)
}

// NewClient instance a new cruptomkt client.
func NewClient(APIKey, secret string) *Client {
	priClient := &httpClient{
		client: &http.Client{},
		key:    APIKey,
		secret: secret,
	}

	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	return &Client{
		PaymentService: PaymentService{priClient, true},
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
