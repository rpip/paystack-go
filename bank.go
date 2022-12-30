package paystack

import "fmt"

// BankService handles operations related to the bank
// For more details see https://developers.paystack.co/v1.0/reference#bank
type BankService service

// Bank represents a Paystack bank
type Bank struct {
	ID        int    `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Name      string `json:"name,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Code      string `json:"code,omitempty"`
	LongCode  string `json:"long_code,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Active    bool   `json:"active,omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty"`
}

// BankList is a list object for banks.
type BankList struct {
	Meta   ListMeta
	Values []Bank `json:"data,omitempty"`
}

type BVNRequest struct {
	BVN           string `json:"bvn,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
}

// BVNResponse represents response from match bvn endpoint
type BVNResponse struct {
	Data struct {
		BVN           string `json:"bvn,omitempty"`
		IsBlacklisted bool   `json:"is_blacklisted,omitempty"`
		AccountNumber bool   `json:"account_number,omitempty"`
		FirstName     bool   `json:"first_name,omitempty"`
		MiddleName    bool   `json:"middle_name,omitempty"`
		LastName      bool   `json:"last_name,omitempty"`
	}
	Meta struct {
		CallsThisMonth int `json:"calls_this_month,omitempty"`
		FreeCallsLeft  int `json:"free_calls_left,omitempty"`
	}
}

// List returns a list of all the banks.
// For more details see https://developers.paystack.co/v1.0/reference#list-banks
func (s *BankService) List() (*BankList, error) {
	banks := &BankList{}
	err := s.client.Call("GET", "/bank", nil, banks)
	return banks, err
}

// MatchBVN docs https://paystack.com/docs/identity-verification/verify-bvn-match/
func (s *BankService) MatchBVN(req *BVNRequest) (*BVNResponse, error) {
	resp := &BVNResponse{}
	err := s.client.Call("POST", "/bvn/match", req, &resp)
	return resp, err
}

// ResolveAccountNumber docs https://developers.paystack.co/v1.0/reference#resolve-account-number
func (s *BankService) ResolveAccountNumber(accountNumber, bankCode string) (Response, error) {
	u := fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", accountNumber, bankCode)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
