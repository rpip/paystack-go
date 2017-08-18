package paystack

import (
	"fmt"
	"net/url"
)

// TransferService handles operations related to the transfer
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer
type TransferService service

// TransactionRequest represents a request to create a transfer.
type TransferRequest struct {
	Source    string  `json:"source,omitempty"`
	Amount    float32 `json:"amount,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	Reason    string  `json:"reason,omitempty"`
	Recipient string  `json:"recipient,omitempty"`
}

type Recipient map[string]interface{}

// Transfer is the resource representing your Paystack transfer.
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
type Transfer struct {
	ID           int       `json:"id,omitempty"`
	CreatedAt    string    `json:"createdAt,omitempty"`
	UpdatedAt    string    `json:"updatedAt,omitempty"`
	Domain       string    `json:"domain,omitempty"`
	Integration  int       `json:"integration,omitempty"`
	Source       string    `json:"source,omitempty"`
	Amount       float32   `json:"amount,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	TransferCode string    `json:"transfer_code,omitempty"`
	Recipient    Recipient `json:"recipient,omitempty"`
	Status       string    `json:"status,omitempty"`
	// confirm types for source_details and failures
	SourceDetails interface{} `json:"source_details,omitempty"`
	Failures      interface{} `json:"failures,omitempty"`
}

type TransferRecipient struct {
	ID            int      `json:"id,omitempty"`
	CreatedAt     string   `json:"createdAt,omitempty"`
	UpdatedAt     string   `json:"updatedAt,omitempty"`
	Type          string   `json:",omitempty"`
	Name          string   `json:"name,omitempty"`
	Metadata      Metadata `json:"metadata,omitempty"`
	AccountNumber string   `json:"account_number,omitempty"`
	BankCode      string   `json:"bank_code,omitempty"`
	Currency      string   `json:"currency,omitempty"`
	Description   string   `json:"description,omitempty"`
}

// BulkTransfer represents a Paystack bulk transfer
type BulkTransfer struct {
	Currency  string           `json:"currency,omitempty"`
	Source    string           `json:"source,omitempty"`
	Transfers []map[string]int `json:"transfers,omitempty"`
}

// TransferList is a list object for transfers.
type TransferList struct {
	Meta   ListMeta
	Values []Transfer `json:"data,omitempty"`
}

// TransferRecipientList is a list object for transfer recipient.
type TransferRecipientList struct {
	Meta   ListMeta
	Values []TransferRecipient `json:"data,omitempty"`
}

// Initialize initiates a new transfer
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
func (s *TransferService) Initialize(req *TransferRequest) (*Transfer, error) {
	u := fmt.Sprintf("/transfer")
	transfer := &Transfer{}
	err := s.client.Call("POST", u, req, transfer)

	return transfer, err
}

// Finalize completes a transfer request
// For more details see https://developers.paystack.co/v1.0/reference#finalize-transfer
func (s *TransferService) Finalize(code, otp string) (*Response, error) {
	u := fmt.Sprintf("/transfer/finalize_transfer")
	req := url.Values{}
	req.Add("transfer_code", code)
	req.Add("otp", otp)

	resp := &Response{}
	err := s.client.Call("POST", u, req, resp)

	return resp, err
}

// Initialize initiates a new bulk transfer
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-transfer
func (s *TransferService) BulkRequest(req *BulkTransfer) (*Response, error) {
	u := fmt.Sprintf("/transfer")
	resp := &Response{}
	err := s.client.Call("POST", u, req, resp)

	return resp, err
}

// Get returns the details of a transfer.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transfer
func (s *TransferService) Get(idCode string) (*Transfer, error) {
	u := fmt.Sprintf("/transfer/%s", idCode)
	transfer := &Transfer{}
	err := s.client.Call("GET", u, nil, transfer)

	return transfer, err
}

// List returns a list of transfers.
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *TransferService) List() (*TransferList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of transfers
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *TransferService) ListN(count, offset int) (*TransferList, error) {
	u := paginateURL("/transfer", count, offset)
	transfers := &TransferList{}
	err := s.client.Call("GET", u, nil, transfers)
	return transfers, err
}

// List returns a list of transfer recipients.
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *TransferService) ListRecipients() (*TransferRecipientList, error) {
	return s.ListRecipientsN(10, 0)
}

// ListN returns a list of transferrecipients
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *TransferService) ListRecipientsN(count, offset int) (*TransferRecipientList, error) {
	u := paginateURL("/transferrecipient", count, offset)
	resp := &TransferRecipientList{}
	err := s.client.Call("GET", u, nil, resp)

	return resp, err
}
