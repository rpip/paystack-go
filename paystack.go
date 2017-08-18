package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	// library version
	version = "0.1.0"

	// defaultHTTPTimeout is the default timeout on the http client
	defaultHTTPTimeout = 60 * time.Second

	// base URL for all Paystack API requests
	baseURL = "https://api.paystack.co"

	// User agent used when communicating with the Paystack API.
	userAgent = "paystack-go/" + version
)

type service struct {
	client *Client
}

// Client manages communication with the Paystack API
type Client struct {
	common service      // Reuse a single struct instead of allocating one for each service on the heap.
	client *http.Client // HTTP client used to communicate with the API.

	// the API Key used to authenticate all Paystack API requests
	key string

	baseURL *url.URL

	// Services supported by the Paystack API.
	// Miscellaneous actions are directly implemented on the Client object
	Customer          *CustomerService
	Transaction       *TransactionService
	SubAccount        *SubAccountService
	Plan              *PlanService
	SubScription      *SubscriptionService
	Page              *PageService
	Settlement        *SettlementService
	Transfer          *TransferService
	TransferRecipient *TransferRecipientService
	BulkCharge        *BulkChargeService
	Charge            *ChargeService
	Bank              *BankService
}

type Metadata map[string]interface{}

// Response represents arbitrary response data
type Response map[string]interface{}

// RequestValues aliased to url.Values as a workaround
type RequestValues url.Values

func (v RequestValues) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3)
	for k, val := range v {
		m[k] = val[0]
	}
	return json.Marshal(m)
}

type ListMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// NewClient creates a new Paystack API client with the given API key
// and HTTP client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(baseURL)
	c := &Client{client: httpClient, key: key, baseURL: u}

	c.common.client = c
	c.Customer = (*CustomerService)(&c.common)
	c.Transaction = (*TransactionService)(&c.common)
	c.SubAccount = (*SubAccountService)(&c.common)
	c.Plan = (*PlanService)(&c.common)
	c.SubScription = (*SubscriptionService)(&c.common)
	c.Page = (*PageService)(&c.common)
	c.Settlement = (*SettlementService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)
	c.TransferRecipient = (*TransferRecipientService)(&c.common)
	c.BulkCharge = (*BulkChargeService)(&c.common)
	c.Charge = (*ChargeService)(&c.common)
	c.Bank = (*BankService)(&c.common)

	return c
}

// s.client.Call("POST", "/v1/plans", PlanRequest{}, &plan)
func (c *Client) Call(method, path string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, _ := c.baseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.key)
	if ua := req.Header.Get("User-Agent"); ua == "" {
		req.Header.Set("User-Agent", userAgent)
	} else {
		req.Header.Set("User-Agent", userAgent+" "+ua)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log error
		return err
	}
	var respMap map[string]interface{}
	json.Unmarshal(respBody, &respMap)
	// fmt.Printf("RESPONSE %s \n %+v", u, respMap)

	if status, _ := respMap["status"].(bool); !status && resp.StatusCode >= 400 {
		err := &Error{
			HTTPStatusCode: resp.StatusCode,
			Message:        respMap["message"].(string),
			URL:            resp.Request.URL,
		}
		if errorDetails, ok := respMap["errors"]; ok {
			err.Errors = errorDetails.(map[string][]interface{})
		}
		return err
	}
	if data, ok := respMap["data"]; ok {
		switch t := respMap["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(respMap, v)
		}
	}
	// response data does not contain data node, return anyways
	return mapstruct(respMap, v)
}

func (c *Client) ResolveCardBIN(bin int) (*Response, error) {
	u := fmt.Sprintf("/decision/bin/%d", bin)
	resp := &Response{}
	err := c.Call("GET", u, nil, resp)

	return resp, err
}

// internals
func paginateURL(path string, count, offset int) string {
	return fmt.Sprintf("%s?perPage=%d&page=%d", path, count, offset)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err = decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

func getTestKey() string {
	key := os.Getenv("PAYSTACK_KEY")

	if len(key) == 0 {
		panic("PAYSTACK_KEY environment variable is not set\n")
	}

	return key
}
