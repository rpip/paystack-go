package paystack

import "fmt"

// BankService handles operations related to the bank
// For more details see https://developers.paystack.co/v1.0/reference#bank
type BankService service

// List returns a list of all the banks.
// For more details see https://developers.paystack.co/v1.0/reference#list-banks
func (s *BankService) List() (*Response, error) {
	resp := &Response{}
	err := s.client.Call("GET", "/bank", nil, resp)

	return resp, err
}

func (s *BankService) ResolveBVN(bvn string) (*Response, error) {
	u := fmt.Sprintf("/bank/resolve_bvn/%s", bvn)
	resp := &Response{}
	err := s.client.Call("GET", u, nil, resp)

	return resp, err
}

func (s *BankService) ResolveAccountNumber(account_number, bank_code int) (*Response, error) {
	u := fmt.Sprintf("/bank/resolve?account_number=%d&bank_code=%d", account_number, bank_code)
	resp := &Response{}
	err := s.client.Call("GET", u, nil, resp)

	return resp, err
}
