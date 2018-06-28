# cryptomkt

<img align="right" width="150" src="gopher.png">


This project implements a Go client library for the [Cryptomkt APIs](https://developers.cryptomkt.com).
This library supports version 1 of Cryptomkt's API.

## Endpoints

### [Public endpoints](https://developers.cryptomkt.com/es/?shell#endpoints-publicos)

#### Markets

- GET /market

   returns a collection of available markets.


```go
package main

import (
	"fmt"

	"github.com/Finciero/go-cryptomkt"
)

func main() {
	// This client expose only public methods.
	cryptomktClient := cryptomkt.NewPublicClient()

	// Request available markets
	response, err := cryptomktClient.GetMarkets()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status) // If everything is OK. then status = success
	fmt.Println(response.Data) // Array of string

	// See public_test.go for more details about usage.
}

```

#### Tickers

- GET /ticker

   returns a collection of active tickers, if the market is present, the specified market ticker is returned.

#### Orders

- GET /book

   returns a collection of active orders.

#### Trades

- GET /trades

   returns a collection of trades made in CryptoMarket.


### [Private endpoints](https://developers.cryptomkt.com/es/?shell#endpoints-autenticados)

### [Cryptocompra](https://developers.cryptomkt.com/es/?shell#cryptocompra)

- POST /payment/new_order

   It allows to create a payment order, delivering QR and urls to pay.

```go
package main

import (
	"fmt"

	"github.com/Finciero/go-cryptomkt"
)

func main() {
	cryptomktKey := "...your key"
	cryptomktSecret := "...your secret"
	cryptomktClient := cryptomkt.NewClient(cryptomktKey, cryptomktSecret)

	// To make a new payment request
	request := &cryptomkt.PaymentRequest{
		Amount:          3000,
		Currency:        "CLP",
		Receiver:        "receiver@email.org",
		ExternalID:      "123456CM",
		NotificationURL: "",
		ErrorURL:        "",
		SuccessURL:      "",
		RefundEmail:     "refund@email.com",
	}
	response, err := cryptomktClient.CreatePayment(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.ID)         // P2023132
	fmt.Println(response.ExternalID) // 123456CM
	fmt.Println(response.Status)     // 0
	fmt.Println(response.QR)         // https://www.cryptomkt.com/invoice/P2023132.png
	fmt.Println(response.PaymentURL) // https://www.cryptomkt.com/invoice/P2023132/xToY232aheSt8F?lang=en
	fmt.Println(response.CreatedAt)  // 2018-06-15T19:44:08.768199
	fmt.Println(response.QR)         // 2018-06-15T19:44:08.768199
}

```

- GET /payment/status

   Returns the status of a payment order

- GET /payment/orders

   Returns the list of generated payment orders


# Tests

Public test use public endpoint of CryptoMarket, and private endpoint use mocked response extracted from [docs](https://developers.cryptomkt.com).

Use this to test the library:

```sh
$ go test
```