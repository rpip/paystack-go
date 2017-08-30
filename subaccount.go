package paystack

import "fmt"

// SubAccountService handles operations related to sub accounts
// For more details see https://developers.paystack.co/v1.0/reference#create-subaccount
type SubAccountService service

// SubAccount is the resource representing your Paystack subaccount.
// For more details see https://developers.paystack.co/v1.0/reference#create-subaccount
type SubAccount struct {
	ID                  int      `json:"id,omitempty"`
	CreatedAt           string   `json:"createdAt,omitempty"`
	UpdatedAt           string   `json:"updatedAt,omitempty"`
	Domain              string   `json:"domain,omitempty"`
	Integration         int      `json:"integration,omitempty"`
	BusinessName        string   `json:"business_name,omitempty"`
	SubAccountCode      string   `json:"subaccount_code,omitempty"`
	Description         string   `json:"description,omitempty"`
	PrimaryContactName  string   `json:"primary_contact_name,omitempty"`
	PrimaryContactEmail string   `json:"primary_contact_email,omitempty"`
	PrimaryContactPhone string   `json:"primary_contact_phone,omitempty"`
	Metadata            Metadata `json:"metadata,omitempty"`
	PercentageCharge    float32  `json:"percentage_charge,omitempty"`
	IsVerified          bool     `json:"is_verified,omitempty"`
	SettlementBank      string   `json:"settlement_bank,omitempty"`
	AccountNumber       string   `json:"account_number,omitempty"`
	SettlementSchedule  string   `json:"settlement_schedule,omitempty"`
	Active              bool     `json:"active,omitempty"`
	Migrate             bool     `json:"migrate,omitempty"`
}

// SubAccountList is a list object for subaccounts.
type SubAccountList struct {
	Meta   ListMeta
	Values []SubAccount `json:"data"`
}

// Create creates a new subaccount
// For more details see https://developers.paystack.co/v1.0/reference#create-subaccount
func (s *SubAccountService) Create(subaccount *SubAccount) (*SubAccount, error) {
	u := fmt.Sprintf("/subaccount")
	acc := &SubAccount{}
	err := s.client.Call("POST", u, subaccount, acc)

	return acc, err
}

// Update updates a subaccount's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-subaccount
// TODO: use ID or slug
func (s *SubAccountService) Update(subaccount *SubAccount) (*SubAccount, error) {
	u := fmt.Sprintf("subaccount/%d", subaccount.ID)
	acc := &SubAccount{}
	err := s.client.Call("PUT", u, subaccount, acc)

	return acc, err
}

// Get returns the details of a subaccount.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-subaccount
// TODO: use ID or slug
func (s *SubAccountService) Get(id int) (*SubAccount, error) {
	u := fmt.Sprintf("/subaccount/%d", id)
	acc := &SubAccount{}
	err := s.client.Call("GET", u, nil, acc)

	return acc, err
}

// List returns a list of subaccounts.
// For more details see https://developers.paystack.co/v1.0/reference#list-subaccounts
func (s *SubAccountService) List() (*SubAccountList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of subaccounts
// For more details see https://developers.paystack.co/v1.0/reference#list-subaccounts
func (s *SubAccountService) ListN(count, offset int) (*SubAccountList, error) {
	u := paginateURL("/subaccount", count, offset)
	acc := &SubAccountList{}
	err := s.client.Call("GET", u, nil, acc)
	return acc, err
}
