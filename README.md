# cryptomkt

<img align="right" width="150" src="gopher.png">


This project implements a Go client library for the [Cryptomkt APIs](https://developers.cryptomkt.com/es/#listado-de-ordenes-de-pago).
This library supports version 1 of Cryptomkt's API.

## Endpoints

### Cryptocompra

- POST /payment/new_order
It allows to create a payment order, delivering QR and urls to pay.

- GET /payment/status
Returns the status of a payment order

- GET /payment/orders
Returns the list of generated payment orders

## Examples