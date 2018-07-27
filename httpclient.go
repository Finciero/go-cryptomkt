package cryptomkt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/bt51/ntpclient"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.cryptomkt.com",
	Path:   "v1",
}

const (
	headerXMktAPIKey        = "X-MKT-APIKEY"
	headerXMktSignature     = "X-MKT-SIGNATURE"
	headerXMktTimestamp     = "X-MKT-TIMESTAMP"
	statusMultiplePayments  = -4
	statusAmountDidNotMatch = -3
	statusConversionFail    = -2
	statusPaymentExpired    = -1
	statusWaitingForPayment = 0
	statusWaitingForBlock   = 1
	statusProcessing        = 2
	statusSuccessfulPayment = 3

	ntpServer = "2.cl.pool.ntp.org"
)

// StatusCodeToText ...
func StatusCodeToText(status int) string {
	switch status {
	case statusMultiplePayments:
		return "multiple-payments"
	case statusAmountDidNotMatch:
		return "invalid-amount"
	case statusConversionFail:
		return "conversion-fail"
	case statusPaymentExpired:
		return "expired"
	case statusWaitingForPayment:
		return "waiting-for-payments"
	case statusWaitingForBlock:
		return "waiting-for-block"
	case statusProcessing:
		return "processing"
	case statusSuccessfulPayment:
		return "success"
	default:
		return "unknown"
	}
}

// httpClient represent a base struct to store Http client configuration
type httpClient struct {
	client  *http.Client
	key     string
	secret  string
	private bool
}

// APIError represents an error of CryptoMKT's REST API.
type APIError struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Error implements error interface.
func (err *APIError) Error() string {
	return fmt.Sprintf("cryptopay: %v", err.Message)
}

func (hc *httpClient) SetPrivate(private bool) {
	hc.private = private
}

func (hc *httpClient) do(req *http.Request, values url.Values) (*http.Response, error) {
	t, err := ntpclient.GetNetworkTime(ntpServer, 123)
	if err != nil {
		return nil, err
	}
	now := t.Unix()

	req.Header.Set("Accept", "application/json")
	if req.Method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if hc.private {
		req.Header.Set(headerXMktAPIKey, hc.key)
		hc.signRequest(req, values, now)
		req.Header.Set(headerXMktTimestamp, fmt.Sprintf("%d", now))
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: %s resquest failed, %v", req.URL, err)
	}

	log.Println(resp.StatusCode)
	// TODO check error response
	switch resp.StatusCode {
	case http.StatusBadRequest:
		var valErr APIError
		if err = unmarshalJSON(resp.Body, &valErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &valErr
	case http.StatusUnauthorized:
		var authErr APIError
		if err = unmarshalJSON(resp.Body, &authErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &authErr
	case http.StatusServiceUnavailable:
		var svcErr APIError
		if err = unmarshalJSON(resp.Body, &svcErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &svcErr
	case http.StatusTooManyRequests:
		var svcErr APIError
		if err = unmarshalJSON(resp.Body, &svcErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &svcErr
	case http.StatusForbidden:
		var svcErr APIError
		if err = unmarshalJSON(resp.Body, &svcErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &svcErr
	case http.StatusNotFound:
		var svcErr APIError
		if err = unmarshalJSON(resp.Body, &svcErr); err != nil {
			return nil, fmt.Errorf("cryptopay: error parsing response, %v", err)
		}
		return nil, &svcErr
	default:
		return resp, nil
	}
}

func (hc *httpClient) get(path string, values url.Values) (*http.Response, error) {
	uri := baseURL.String() + path

	if values == nil {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			return nil, err
		}
		return hc.do(req, values)
	}

	req, err := http.NewRequest(http.MethodGet, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	return hc.do(req, values)
}

func (hc *httpClient) postForm(path string, values url.Values) (*http.Response, error) {
	url := baseURL.String() + path

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	return hc.do(req, values)
}

func (hc *httpClient) signRequest(req *http.Request, values url.Values, timestamp int64) {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("%d", timestamp))
	buff.WriteString(req.URL.Path)

	switch req.URL.Path {
	case "/v1/orders/cancel":
		keys := make([]string, 0)
		for k := range values {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buff.WriteString(values[k][0])
		}
		break
	case "/v1/payment/new_order":
		if values != nil {
			keys := make([]string, 0)
			for k := range values {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				buff.WriteString(values[k][0])
			}
		}
		break
	}

	sig := hmac.New(sha512.New384, []byte(hc.secret))
	sig.Write(buff.Bytes())
	sign := hex.EncodeToString(sig.Sum(nil))
	buff.Reset()

	req.Header.Set(headerXMktSignature, sign)
}

func unmarshalJSON(r io.ReadCloser, v interface{}) error {
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return json.Unmarshal(body, v)
}

// SpecialInt ...
type SpecialInt int

// UnmarshalJSON ...
func (si *SpecialInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(si))
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "null" || s == "" {
		*si = SpecialInt(0)
		return nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*si = SpecialInt(i)
	return nil
}
