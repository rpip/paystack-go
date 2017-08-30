package paystack

import "fmt"

// TransactionService handles operations related to transactions
// For more details see https://developers.paystack.co/v1.0/reference#create-transaction
type TransactionService service

// TransactionList is a list object for transactions.
type TransactionList struct {
	Meta   ListMeta
	Values []Transaction `json:"data"`
}

// TransactionRequest represents a request to start a transaction.
type TransactionRequest struct {
	CallbackURL       string   `json:"callback_url,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	AuthorizationCode string   `json:"authorization_code,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Amount            float32  `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      int      `json:"invoice_limit,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
	SubAccount        string   `json:"subaccount,omitempty"`
	TransactionCharge int      `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
	Channels          []string `json:"channels,omitempty"`
}

// AuthorizationRequest represents a request to enable/revoke an authorization
type AuthorizationRequest struct {
	Reference         string   `json:"reference,omitempty"`
	AuthorizationCode string   `json:"authorization_code,omitempty"`
	Amount            int      `json:"amount,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Email             string   `json:"email,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
}

// Transaction is the resource representing your Paystack transaction.
// For more details see https://developers.paystack.co/v1.0/reference#initialize-a-transaction
type Transaction struct {
	ID              int                    `json:"id,omitempty"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
	Domain          string                 `json:"domain,omitempty"`
	Metadata        string                 `json:"metadata,omitempty"` //TODO: why is transaction metadata a string?
	Status          string                 `json:"status,omitempty"`
	Reference       string                 `json:"reference,omitempty"`
	Amount          float32                `json:"amount,omitempty"`
	Message         string                 `json:"message,omitempty"`
	GatewayResponse string                 `json:"gateway_response,omitempty"`
	PaidAt          string                 `json:"piad_at,omitempty"`
	Channel         string                 `json:"channel,omitempty"`
	Currency        string                 `json:"currency,omitempty"`
	IPAddress       string                 `json:"ip_address,omitempty"`
	Log             map[string]interface{} `json:"log,omitempty"` // TODO: same as timeline?
	Fees            int                    `json:"int,omitempty"`
	FeesSplit       string                 `json:"fees_split,omitempty"` // TODO: confirm data type
	Customer        Customer               `json:"customer,omitempty"`
	Authorization   Authorization          `json:"authorization,omitempty"`
	Plan            Plan                   `json:"plan,omitempty"`
	SubAccount      SubAccount             `json:"sub_account,omitempty"`
}

// Authorization represents Paystack authorization object
type Authorization struct {
	AuthorizationCode string `json:"authorization_code,omitempty"`
	Bin               string `json:"bin,omitempty"`
	Last4             string `json:"last4,omitempty"`
	ExpMonth          string `json:"exp_month,omitempty"`
	ExpYear           string `json:"exp_year,omitempty"`
	Channel           string `json:"channel,omitempty"`
	CardType          string `json:"card_type,omitempty"`
	Bank              string `json:"bank,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
	Brand             string `json:"brand,omitempty"`
	Resusable         bool   `json:"reusable,omitempty"`
	Signature         string `json:"signature,omitempty"`
}

// TransactionTimeline represents a timeline of events in a transaction session
type TransactionTimeline struct {
	TimeSpent      int                      `json:"time_spent,omitempty"`
	Attempts       int                      `json:"attempts,omitempty"`
	Authentication string                   `json:"authentication,omitempty"` // TODO: confirm type
	Errors         int                      `json:"errors,omitempty"`
	Success        bool                     `json:"success,omitempty"`
	Mobile         bool                     `json:"mobile,omitempty"`
	Input          []string                 `json:"input,omitempty"` // TODO: confirm type
	Channel        string                   `json:"channel,omitempty"`
	History        []map[string]interface{} `json:"history,omitempty"`
}

// Initialize initiates a transaction process
// For more details see https://developers.paystack.co/v1.0/reference#initialize-a-transaction
func (s *TransactionService) Initialize(txn *TransactionRequest) (Response, error) {
	u := fmt.Sprintf("/transaction/initialize")
	resp := Response{}
	err := s.client.Call("POST", u, txn, &resp)
	return resp, err
}

// Verify checks that transaction with the given reference exists
// For more details see https://api.paystack.co/transaction/verify/reference
func (s *TransactionService) Verify(reference string) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/verify/%s", reference)
	txn := &Transaction{}
	err := s.client.Call("GET", u, nil, txn)
	return txn, err
}

// List returns a list of transactions.
// For more details see https://developers.paystack.co/v1.0/reference#list-transactions
func (s *TransactionService) List() (*TransactionList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of transactions
// For more details see https://developers.paystack.co/v1.0/reference#list-transactions
func (s *TransactionService) ListN(count, offset int) (*TransactionList, error) {
	u := paginateURL("/transaction", count, offset)
	txns := &TransactionList{}
	err := s.client.Call("GET", u, nil, txns)
	return txns, err
}

// Get returns the details of a transaction.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transaction
func (s *TransactionService) Get(id int) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/%d", id)
	txn := &Transaction{}
	err := s.client.Call("GET", u, nil, txn)
	return txn, err
}

// ChargeAuthorization is for charging all  authorizations marked as reusable whenever you need to recieve payments.
// For more details see https://developers.paystack.co/v1.0/reference#charge-authorization
func (s *TransactionService) ChargeAuthorization(req *TransactionRequest) (*Transaction, error) {
	txn := &Transaction{}
	err := s.client.Call("POST", "/transaction/charge_authorization", req, txn)
	return txn, err
}

// Timeline fetches the transaction timeline. Reference can be ID or transaction reference
// For more details see https://developers.paystack.co/v1.0/reference#view-transaction-timeline
func (s *TransactionService) Timeline(reference string) (*TransactionTimeline, error) {
	u := fmt.Sprintf("/transaction/timeline/%s", reference)
	timeline := &TransactionTimeline{}
	err := s.client.Call("GET", u, nil, timeline)
	return timeline, err
}

// Totals returns total amount received on your account
// For more details see https://developers.paystack.co/v1.0/reference#transaction-totals
func (s *TransactionService) Totals() (Response, error) {
	u := fmt.Sprintf("/transaction/totals")
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}

// Export exports transactions to a downloadable file and returns a link to the file
// For more details see https://developers.paystack.co/v1.0/reference#export-transactions
func (s *TransactionService) Export(params RequestValues) (Response, error) {
	u := fmt.Sprintf("/transaction/export")
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}

// ReAuthorize requests reauthorization
// For more details see https://developers.paystack.co/v1.0/reference#request-reauthorization
func (s *TransactionService) ReAuthorize(req AuthorizationRequest) (Response, error) {
	u := fmt.Sprintf("/transaction/request_reauthorization")
	resp := Response{}
	err := s.client.Call("POST", u, nil, &resp)
	return resp, err
}

// CheckAuthorization checks authorization
// For more details see https://developers.paystack.co/v1.0/reference#check-authorization
func (s *TransactionService) CheckAuthorization(req AuthorizationRequest) (Response, error) {
	u := fmt.Sprintf("/transaction/check_reauthorization")
	resp := Response{}
	err := s.client.Call("POST", u, nil, &resp)
	return resp, err
}
