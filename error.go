package paystack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIError struct {
	Message        string                `json:"message,omitempty"`
	HTTPStatusCode int                   `json:"code,omitempty"`
	Details        PaystackErrorResponse `json:"details,omitempty"`
	URL            *url.URL              `json:"url,omitempty"`
	Header         http.Header           `json:"header,omitempty"`
}

// APIError supports the error interface
func (aerr *APIError) Error() string {
	ret, _ := json.Marshal(aerr)
	return string(ret)
}

// PaystackErrorResponse represents an error response from the Paystack API server
type PaystackErrorResponse struct {
	Status  bool                   `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

func newAPIError(resp *http.Response) *APIError {
	p, _ := ioutil.ReadAll(resp.Body)

	var paystackErrorResp PaystackErrorResponse
	_ = json.Unmarshal(p, &paystackErrorResp)
	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		Header:         resp.Header,
		Details:        paystackErrorResp,
		URL:            resp.Request.URL,
	}
}
