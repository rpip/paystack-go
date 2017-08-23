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

// TODO: return meta data from BVN response
func (s *BankService) ResolveBVN(bvn int) (*BVNResponse, error) {
	u := fmt.Sprintf("/bank/resolve_bvn/%d", bvn)
	resp := &BVNResponse{}
	err := s.client.Call("GET", u, nil, resp)
	return resp, err
}

func (s *BankService) ResolveAccountNumber(account_number, bank_code string) (Response, error) {
	u := fmt.Sprintf("/bank/resolve?account_number=%d&bank_code=%d", account_number, bank_code)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}
