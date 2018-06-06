# cryptomkt

<img align="right" width="150" src="gopher.png">


This project implements a Go client library for the [Cryptomkt APIs](https://developers.cryptomkt.com).
This library supports version 1 of Cryptomkt's API.

## Endpoints

### [Cryptocompra](https://developers.cryptomkt.com/es/?shell#cryptocompra)

- POST /payment/new_order

   It allows to create a payment order, delivering QR and urls to pay.

- GET /payment/status


   Returns the status of a payment order

- GET /payment/orders

   Returns the list of generated payment orders

## Examples