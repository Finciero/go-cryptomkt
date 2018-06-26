package cryptomkt

import (
	"net/http"
	"testing"
)

func Test_GetMarkets(t *testing.T) {
	ps := &PublicService{
		client: &httpClient{
			client: &http.Client{},
		},
		Private: false,
	}

	mr, err := ps.GetMarkets()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if mr.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mr.Status)
	}
}

func Test_GetTicker(t *testing.T) {
	ps := &PublicService{
		client: &httpClient{
			client: &http.Client{},
		},
		Private: false,
	}

	tr, err := ps.GetTicker("ETHARS")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if tr.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, tr.Status)
	}
}

func Test_GetBooks(t *testing.T) {
	ps := &PublicService{
		client: &httpClient{
			client: &http.Client{},
		},
		Private: false,
	}

	opts := &BooksOptions{
		Market: "ETHCLP",
		Type:   "buy",
		Page:   1,
	}
	br, err := ps.GetOrdersBook(opts)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if br.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, br.Status)
	}
}

func Test_GetTrades(t *testing.T) {
	ps := &PublicService{
		client: &httpClient{
			client: &http.Client{},
		},
		Private: false,
	}

	opts := &TradesOptions{
		Market:    "ETHCLP",
		StartDate: "2017-05-20",
		EndDate:   "2017-05-30",
		Page:      2,
	}
	tr, err := ps.GetTrades(opts)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "success"
	if tr.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, tr.Status)
	}
}
