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
	"net/http"
	"net/url"
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
	statusMultiplePayments  = "-4"
	statusAmountDidNotMatch = "-3"
	statusConversionFail    = "-2"
	statusPaymentExpired    = "-1"
	statusWaitingForPayment = "0"
	statusWaitingForBlock   = "1"
	statusProcessing        = "2"
	statusSuccessfulPayment = "3"

	ntpServer = "2.cl.pool.ntp.org"
)

// httpClient represent a base struct to store Http client configuration
type httpClient struct {
	client *http.Client
	key    string
	secret string
}

// APIError represents an error of CryptoMKT's REST API.
type APIError struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Error implements error interface.
func (err *APIError) Error() string {
	return fmt.Sprintf("cryptopay: unauthorized request, %v", err.Message)
}

func (hc *httpClient) do(req *http.Request, values url.Values) (*http.Response, error) {
	t, err := ntpclient.GetNetworkTime(ntpServer, 123)
	if err != nil {
		return nil, err
	}
	now := t.Unix()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(headerXMktAPIKey, hc.key)
	hc.signRequest(req, values, now)
	req.Header.Set(headerXMktTimestamp, fmt.Sprintf("%d", now))

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: %s resquest failed, %v", req.URL, err)
	}

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
	default:
		return resp, nil
	}
}

func (hc *httpClient) get(path string, values url.Values) (*http.Response, error) {
	uri := baseURL.String() + path
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return hc.do(req, values)
}

func (hc *httpClient) postForm(path string, values url.Values) (*http.Response, error) {
	uri := baseURL.String() + path
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	return hc.do(req, values)
}

var percentEncode = strings.NewReplacer(
	"+", "%20",
	"*", "%2A",
	"%7A", "~",
)

func (hc *httpClient) signRequest(req *http.Request, values url.Values, timestamp int64) {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("%d", timestamp))
	buff.WriteString(req.URL.Path)
	if values != nil {
		for _, value := range values {
			buff.WriteString(value[0])
		}
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

	fmt.Printf("\n\n%s\n\n", body)
	return json.Unmarshal(body, v)
}
