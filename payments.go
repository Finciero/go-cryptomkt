package cryptomkt

import (
	"errors"
	"fmt"
	"net/url"
)

// PaymentService ...
type PaymentService struct {
	client *httpClient
}

func checkStatus(status string) error {
	switch status {
	case statusMultiplePayments:
		return errors.New("cryptopay: Multple payments")
	case statusAmountDidNotMatch:
		return errors.New("cryptopay: Amount didn't match")
	case statusConversionFail:
		return errors.New("cryptopay: Convertion failed")
	case statusPaymentExpired:
		return errors.New("cryptopay: Payment expired")
	default:
		return nil
	}
}

// CreatePayment ...
func (ps *PaymentService) CreatePayment(p *PaymentRequest) (*PaymentResponse, error) {
	resp, err := ps.client.postForm("/payment/new_order", p.Params())
	if err != nil {
		return nil, err
	}

	var r Response
	if err := unmarshalJSON(resp.Body, &r); err != nil {
		return nil, err
	}

	if err := checkStatus(r.Response.Status); err != nil {
		return nil, err
	}

	return r.Response, nil
}

// PaymentStatus ...
func (ps *PaymentService) PaymentStatus(id string) (*PaymentResponse, error) {
	p := url.Values{
		"id": {id},
	}
	resp, err := ps.client.get("/payment/status?", p)
	if err != nil {
		return nil, err
	}

	var r Response
	if err := unmarshalJSON(resp.Body, &r); err != nil {
		return nil, err
	}

	if err := checkStatus(r.Response.Status); err != nil {
		return nil, err
	}

	return r.Response, nil
}

// Response ...
type Response struct {
	Status     string           `json:"status"`
	Response   *PaymentResponse `json:"data"`
	Pagination interface{}      `json:"pagination"`
}

// PaymentRequest represents the payment form requires by khipu to make a payment POST
type PaymentRequest struct {
	// Monto a cobrar de la orden de pago. CLP no soporta decimales.
	Amount int64 `json:"to_receive"`
	// Tipo de moneda con la cual recibirá el pago
	Currency string `json:"to_receive_currency"`
	// Email del usuario o comercio que recibirá el pago. Debe estar registrado en CryptoMarket.
	Receiver string `json:"payment_receiver"`
	// ID externo. Permite asociar orden interna de comercio con orden de pago. Max. 64 caracteres.
	ExternalID string `json:"external_id"`
	// Url a la cual se notificarán los cambios de estado de la orden. Max. 256 caracteres.
	NotificationURL string `json:"callback_url"`
	// Url a la cual se rediccionará en caso de error. Max. 256 caracteres.
	ErrorURL string `json:"error_url"`
	// Url a la cual se rediccionará en caso de éxito. Max. 256 caracteres.
	SuccessURL string `json:"success_url"`
	// Correo electrónico de contacto para coordinar reembolsos
	RefundEmail string `json:"refund_email"`
}

// Params returns a map used to sign the requests
func (p *PaymentRequest) Params() url.Values {
	form := url.Values{
		"to_receive":          {fmt.Sprintf("%d", p.Amount)},
		"to_receive_currency": {p.Currency},
		"payment_receiver":    {p.Receiver},
		"external_id":         {p.ExternalID},
		"callback_url":        {p.NotificationURL},
		"error_url":           {p.ErrorURL},
		"success_url":         {p.SuccessURL},
		"refund_email":        {p.RefundEmail},
	}

	return form
}

// PaymentResponse ...
type PaymentResponse struct {
	// ID interno de la orden de pago
	ID string `json:"id"`
	// ID externo
	ExternalID string `json:"external_id"`
	// Estado de la orden de pago
	Status string `json:"status"`
	// URL de la imagen QR de la orden de pago
	QR string `json:"qr"`
	// URL de voucher de ordern de pago
	PaymentURL string `json:"payment_url"`
	// Fecha de creación de la orden de pago
	CreatedAt string `json:"created_at"`
	// Fecha de actualización de la orden de pago
	UpdatedAt string `json:"updated_at"`
}
