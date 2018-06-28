package cryptomkt

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetActiveOrders(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getActiveOrdersResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	moo := &MarketOrderOptions{
		Market: "ETHCL",
		Page:   0,
		Limit:  10,
	}
	mor, err := ps.GetActiveOrders(moo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
		return
	}

	if mor.Pagination == nil {
		t.Errorf("Expected Pagination not be nil")
		return
	}

	if mor.Data == nil {
		t.Errorf("Expected Data not be nil")
		return
	}

	expectedLength := 2
	actualLength := len(mor.Data)
	if actualLength != expectedLength {
		t.Errorf("Expected Data length to be %d, got %d", expectedLength, actualLength)
		return
	}

	expectedMaket := "ETHCLP"
	actualMarket := mor.Data[0].Market
	if actualMarket != expectedMaket {
		t.Errorf("Expected market to be %s, got %s", expectedMaket, actualMarket)
		return
	}
}

func Test_GetExecutedOrders(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getExecutedOrdersResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	moo := &MarketOrderOptions{
		Market: "ETHCL",
		Page:   0,
		Limit:  10,
	}
	mor, err := ps.GetExecutedOrders(moo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
		return
	}

	if mor.Pagination == nil {
		t.Errorf("Expected Pagination not be nil")
		return
	}

	if mor.Data == nil {
		t.Errorf("Expected Data not be nil")
		return
	}

	expectedLength := 2
	actualLength := len(mor.Data)
	if actualLength != expectedLength {
		t.Errorf("Expected Data length to be %d, got %d", expectedLength, actualLength)
		return
	}

	expectedMaket := "ETHCLP"
	actualMarket := mor.Data[0].Market
	if actualMarket != expectedMaket {
		t.Errorf("Expected market to be %s, got %s", expectedMaket, actualMarket)
		return
	}
}

func Test_CreateOrder(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getCreateOrderResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	mor := &MarketOrderRequest{
		Market: "ethclp",
		Amount: 0.3,
		Price:  10000,
		Type:   "buy",
	}
	morr, err := ps.CreateOrder(mor)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expectedStatus := "success"
	actualStatus := morr.Status
	if actualStatus != expectedStatus {
		t.Errorf("Expected status %s, got %s", expectedStatus, actualStatus)
		return
	}

	if morr.Data == nil {
		t.Errorf("Expected Data not be nil")
		return
	}

	expectedMaket := "ETHCLP"
	actualMarket := morr.Data.Market
	if actualMarket != expectedMaket {
		t.Errorf("Expected market to be %s, got %s", expectedMaket, actualMarket)
		return
	}
}

func Test_GetOrderStatus(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getStatusOrderResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	oso := &OrderStatusOption{
		ID: "M103975",
	}
	mor, err := ps.GetOrderStatus(oso)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
		return
	}

	if mor.Data == nil {
		t.Errorf("Expected Data not be nil")
		return
	}

	expectedMaket := "ETHCLP"
	actualMarket := mor.Data.Market
	if actualMarket != expectedMaket {
		t.Errorf("Expected market to be %s, got %s", expectedMaket, actualMarket)
		return
	}
}

func Test_CancelOrder(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getCancelOrderResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	cor := &CancelOrderRequest{
		ID: "M103975",
	}
	mor, err := ps.CancelOrder(cor)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := "success"
	if mor.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, mor.Status)
		return
	}

	if mor.Data == nil {
		t.Errorf("Expected Data not be nil")
		return
	}

	expectedMaket := "ETHCLP"
	actualMarket := mor.Data.Market
	if actualMarket != expectedMaket {
		t.Errorf("Expected market to be %s, got %s", expectedMaket, actualMarket)
		return
	}
}

func Test_GetBalance(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(getBalanceResponse)
	})
	httpCli, teardown := testingHTTPClient(h)
	defer teardown()
	ps := &PrivateService{
		client: &httpClient{
			client:  httpCli,
			key:     "some-key",
			secret:  "some-secret",
			private: true,
		},
		Private: true,
	}

	bal, err := ps.GetBalance()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := "success"
	if bal.Status != "success" {
		t.Errorf("Expected status %s, got %s", expected, bal.Status)
		return
	}
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}

var getBalanceResponse = []byte(`
	{
		"status": "success",
		"data": [
		   {
			  "available": "120347",
			  "wallet": "CLP",
			  "balance": "120347"
		   },
		   {
			  "available": "10.3399",
			  "wallet": "ETH",
			  "balance": "11.3399"
		   }
		]
	 }
`)

var getActiveOrdersResponse = []byte(`
	{
		"status": "success",
		"pagination": {
		   "previous": "null",
		   "limit": 20,
		   "page": 0,
		   "next": "null"
		},
		"data": [
		   {
			  "status": "active",
			  "created_at": "2017-09-01T14:01:56.887272",
			  "amount": {
				 "original": "1.4044",
				 "remaining": "1.4044"
			  },
			  "execution_price": null,
			  "price": "7120",
			  "type": "buy",
			  "id": "M103966",
			  "market": "ETHCLP",
			  "updated_at": "2017-09-01T14:01:56.887272"
		   },
		   {
			  "status": "active",
			  "created_at": "2017-09-01T14:02:36.386967",
			  "amount": {
				 "original": "1.25",
				 "remaining": "1.25"
			  },
			  "execution_price": null,
			  "price": "8000",
			  "type": "buy",
			  "id": "M103967",
			  "market": "ETHCLP",
			  "updated_at": "2017-09-01T14:02:36.386967"
		   }
		]
	 }
`)

var getExecutedOrdersResponse = []byte(`
	{
		"status": "success",
		"pagination": {
		   "previous": "null",
		   "limit": 20,
		   "page": 0,
		   "next": "null"
		},
		"data": [
		   {
			  "status": "executed",
			  "created_at": "2017-08-31T21:37:42.282102",
			  "amount": {
				 "executed": "0.6",
				 "original": "3.75"
			  },
			  "execution_price": "8000",
			  "executed_at": "2017-08-31T22:01:19.481403",
			  "price": "8000",
			  "type": "buy",
			  "id": "M103959",
			  "market": "ETHCLP"
		   },
		   {
			  "status": "executed",
			  "created_at": "2017-08-31T21:37:42.282102",
			  "amount": {
				 "executed": "0.5",
				 "original": "3.75"
			  },
			  "execution_price": "8000",
			  "executed_at": "2017-08-31T22:00:13.805482",
			  "price": "8000",
			  "type": "buy",
			  "id": "M103959",
			  "market": "ETHCLP"
		   },
		   {
			  "status": "executed",
			  "created_at": "2016-11-26T23:27:54.502024",
			  "amount": {
				 "executed": "1.5772",
				 "original": "1.5772"
			  },
			  "execution_price": "6340",
			  "executed_at": "2017-01-02T22:56:03.897534",
			  "price": "6340",
			  "type": "buy",
			  "id": "M103260",
			  "market": "ETHCLP"
		   }
		]
	 }
`)

var getCreateOrderResponse = []byte(`
	{
		"status": "success",
		"data": {
		   "status": "executed",
		   "created_at": "2017-09-01T19:35:26.641136",
		   "amount": {
			  "executed": "0.3",
			  "original": "0.3"
		   },
		   "avg_execution_price": "30000",
		   "price": "10000",
		   "type": "buy",
		   "id": "M103975",
		   "market": "ETHCLP",
		   "updated_at": "2017-09-01T19:35:26.688106"
		}
	 }
	`)

var getStatusOrderResponse = []byte(`
	{
		"status": "success",
		"data": {
		   "status": "active",
		   "created_at": "2017-09-01T14:01:56.887272",
		   "amount": {
			  "executed": "0",
			  "original": "1.4044"
		   },
		   "avg_execution_price": "0",
		   "price": "7120",
		   "type": "buy",
		   "id": "M103966",
		   "market": "ETHCLP",
		   "updated_at": "2017-09-01T14:01:56.887272"
		}
	 }
`)

var getCancelOrderResponse = []byte(`
	{
		"status": "success",
		"data": {
		   "status": "cancelled",
		   "created_at": "2017-09-01T14:02:36.386967",
		   "amount": {
			  "executed": "0",
			  "original": "1.25"
		   },
		   "avg_execution_price": "0",
		   "price": "8000",
		   "type": "buy",
		   "id": "M103967",
		   "market": "ETHCLP",
		   "updated_at": "2017-09-01T14:02:36.386967"
		}
	 }
`)
