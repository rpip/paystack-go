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

// BVNResponse represents response from resolve_bvn endpoint
type BVNResponse struct {
	Meta struct {
		CallsThisMonth int `json:"calls_this_month,omitempty"`
		FreeCallsLeft  int `json:"free_calls_left,omitempty"`
	}
	BVN string
}

// List returns a list of all the banks.
// For more details see https://developers.paystack.co/v1.0/reference#list-banks
func (s *BankService) List() (*BankList, error) {
	banks := &BankList{}
	err := s.client.Call("GET", "/bank", nil, banks)
	return banks, err
}

// ResolveBVN docs https://developers.paystack.co/v1.0/reference#resolve-bvn
func (s *BankService) ResolveBVN(bvn int) (*BVNResponse, error) {
	u := fmt.Sprintf("/bank/resolve_bvn/%d", bvn)
	resp := &BVNResponse{}
	err := s.client.Call("GET", u, nil, resp)
	return resp, err
}

// ResolveAccountNumber docs https://developers.paystack.co/v1.0/reference#resolve-account-number
func (s *BankService) ResolveAccountNumber(accountNumber, bankCode string) (Response, error) {
	u := fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", accountNumber, bankCode)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
