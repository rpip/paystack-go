package paystack

import (
	"fmt"
	"net/url"
)

// TransferService handles operations related to the transfer
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer
type TransferService service

// TransferRequest represents a request to create a transfer.
type TransferRequest struct {
	Source    string  `json:"source,omitempty"`
	Amount    float32 `json:"amount,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	Reason    string  `json:"reason,omitempty"`
	Recipient string  `json:"recipient,omitempty"`
}

// Transfer is the resource representing your Paystack transfer.
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
type Transfer struct {
	ID           int     `json:"id,omitempty"`
	CreatedAt    string  `json:"createdAt,omitempty"`
	UpdatedAt    string  `json:"updatedAt,omitempty"`
	Domain       string  `json:"domain,omitempty"`
	Integration  int     `json:"integration,omitempty"`
	Source       string  `json:"source,omitempty"`
	Amount       float32 `json:"amount,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Reason       string  `json:"reason,omitempty"`
	TransferCode string  `json:"transfer_code,omitempty"`
	// Initiate returns recipient ID as recipient value, Fetch returns recipient object
	Recipient interface{} `json:"recipient,omitempty"`
	Status    string      `json:"status,omitempty"`
	// confirm types for source_details and failures
	SourceDetails interface{} `json:"source_details,omitempty"`
	Failures      interface{} `json:"failures,omitempty"`
	TransferredAt string      `json:"transferred_at,omitempty"`
	TitanCode     string      `json:"titan_code,omitempty"`
}

// TransferRecipient represents a Paystack transfer recipient
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer-recipient
type TransferRecipient struct {
	ID            int                    `json:"id,omitempty"`
	CreatedAt     string                 `json:"createdAt,omitempty"`
	UpdatedAt     string                 `json:"updatedAt,omitempty"`
	Type          string                 `json:",omitempty"`
	Name          string                 `json:"name,omitempty"`
	Metadata      Metadata               `json:"metadata,omitempty"`
	AccountNumber string                 `json:"account_number,omitempty"`
	BankCode      string                 `json:"bank_code,omitempty"`
	Currency      string                 `json:"currency,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Active        bool                   `json:"active,omitempty"`
	Details       map[string]interface{} `json:"details,omitempty"`
	Domain        string                 `json:"domain,omitempty"`
	RecipientCode string                 `json:"recipient_code,omitempty"`
}

// BulkTransfer represents a Paystack bulk transfer
// You need to disable the Transfers OTP requirement to use this endpoint
type BulkTransfer struct {
	Currency  string                   `json:"currency,omitempty"`
	Source    string                   `json:"source,omitempty"`
	Transfers []map[string]interface{} `json:"transfers,omitempty"`
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

// Initiate initiates a new transfer
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
func (s *TransferService) Initiate(req *TransferRequest) (*Transfer, error) {
	transfer := &Transfer{}
	err := s.client.Call("POST", "/transfer", req, transfer)
	return transfer, err
}

// Finalize completes a transfer request
// For more details see https://developers.paystack.co/v1.0/reference#finalize-transfer
func (s *TransferService) Finalize(code, otp string) (Response, error) {
	u := fmt.Sprintf("/transfer/finalize_transfer")
	req := url.Values{}
	req.Add("transfer_code", code)
	req.Add("otp", otp)
	resp := Response{}
	err := s.client.Call("POST", u, req, &resp)
	return resp, err
}

// MakeBulkTransfer initiates a new bulk transfer request
// You need to disable the Transfers OTP requirement to use this endpoint
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-transfer
func (s *TransferService) MakeBulkTransfer(req *BulkTransfer) (Response, error) {
	u := fmt.Sprintf("/transfer")
	resp := Response{}
	err := s.client.Call("POST", u, req, &resp)
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

// ResendOTP generates a new OTP and sends to customer in the event they are having trouble receiving one.
// For more details see https://developers.paystack.co/v1.0/reference#resend-otp-for-transfer
func (s *TransferService) ResendOTP(transferCode, reason string) (Response, error) {
	data := url.Values{}
	data.Add("transfer_code", transferCode)
	data.Add("reason", reason)
	resp := Response{}
	err := s.client.Call("POST", "/transfer/resend_otp", data, &resp)
	return resp, err
}

// EnableOTP enables OTP requirement for Transfers
// In the event that a customer wants to stop being able to complete
// transfers programmatically, this endpoint helps turn OTP requirement back on.
// No arguments required.
func (s *TransferService) EnableOTP() (Response, error) {
	resp := Response{}
	err := s.client.Call("POST", "/transfer/enable_otp", nil, &resp)
	return resp, err
}

// DisableOTP disables OTP requirement for Transfers
// In the event that you want to be able to complete transfers
// programmatically without use of OTPs, this endpoint helps disable that….
// with an OTP. No arguments required. You will get an OTP.
func (s *TransferService) DisableOTP() (Response, error) {
	resp := Response{}
	err := s.client.Call("POST", "/transfer/disable_otp", nil, &resp)
	return resp, err
}

// FinalizeOTPDisable finalizes disabling of OTP requirement for Transfers
// For more details see https://developers.paystack.co/v1.0/reference#finalize-disabling-of-otp-requirement-for-transfers
func (s *TransferService) FinalizeOTPDisable(otp string) (Response, error) {
	data := url.Values{}
	data.Add("otp", otp)
	resp := Response{}
	err := s.client.Call("POST", "/transfer/disable_otp_finalize", data, &resp)
	return resp, err
}

// CreateRecipient creates a new transfer recipient
// For more details see https://developers.paystack.co/v1.0/reference#create-transferrecipient
func (s *TransferService) CreateRecipient(recipient *TransferRecipient) (*TransferRecipient, error) {
	recipient1 := &TransferRecipient{}
	err := s.client.Call("POST", "/transferrecipient", recipient, recipient1)
	return recipient1, err
}

// ListRecipients returns a list of transfer recipients.
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *TransferService) ListRecipients() (*TransferRecipientList, error) {
	return s.ListRecipientsN(10, 0)
}

// ListRecipientsN returns a list of transfer recipients
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *TransferService) ListRecipientsN(count, offset int) (*TransferRecipientList, error) {
	u := paginateURL("/transferrecipient", count, offset)
	resp := &TransferRecipientList{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
