package cryptomkt

import (
	"net/http"
	"os"
	"testing"
)

func Test_GetActiveOrders(t *testing.T) {
	key := os.Getenv("CRYPTOMKT_KEY")
	if key == "" {
		t.Errorf("CRYPTOMKT_KEY need to be set.")
		return
	}

	secret := os.Getenv("CRYPTOMKT_SECRET")
	if key == "" {
		t.Errorf("CRYPTOMKT_SECRET need to be set.")
		return
	}
	ps := &PrivateService{
		client: &httpClient{
			client:  &http.Client{},
			key:     key,
			secret:  secret,
			private: true,
		},
		Private: true,
	}

	moo := &MarketOrderOptions{
		Market: "ETHARS",
		Page:   1,
		Limit:  10,
	}
	mor, err := ps.GetActiveOrders(moo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
	}
}

func Test_GetExecutedOrders(t *testing.T) {
	key := os.Getenv("CRYPTOMKT_KEY")
	if key == "" {
		t.Errorf("CRYPTOMKT_KEY need to be set.")
		return
	}

	secret := os.Getenv("CRYPTOMKT_SECRET")
	if key == "" {
		t.Errorf("CRYPTOMKT_SECRET need to be set.")
		return
	}
	ps := &PrivateService{
		client: &httpClient{
			client:  &http.Client{},
			key:     key,
			secret:  secret,
			private: true,
		},
		Private: true,
	}

	moo := &MarketOrderOptions{
		Market: "ETHARS",
		Page:   0,
		Limit:  10,
	}
	mor, err := ps.GetExecutedOrders(moo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
	}
}

func Test_GetBalance(t *testing.T) {
	key := os.Getenv("CRYPTOMKT_KEY")
	if key == "" {
		t.Errorf("CRYPTOMKT_KEY need to be set.")
		return
	}

	secret := os.Getenv("CRYPTOMKT_SECRET")
	if key == "" {
		t.Errorf("CRYPTOMKT_SECRET need to be set.")
		return
	}
	ps := &PrivateService{
		client: &httpClient{
			client:  &http.Client{},
			key:     key,
			secret:  secret,
			private: true,
		},
		Private: true,
	}

	bal, err := ps.GetBalance()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if bal.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, bal.Status)
	}
}
