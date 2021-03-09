package paystack

import (
	"fmt"
	"net/url"
)

// CustomerService handles operations related to the customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
type CustomerService service

// Customer is the resource representing your Paystack customer.
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
type Customer struct {
	ID             int            `json:"id,omitempty"`
	CreatedAt      string         `json:"createdAt,omitempty"`
	UpdatedAt      string         `json:"updatedAt,omitempty"`
	Domain         string         `json:"domain,omitempty"`
	Integration    int            `json:"integration,omitempty"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	Email          string         `json:"email,omitempty"`
	Phone          string         `json:"phone,omitempty"`
	Metadata       Metadata       `json:"metadata,omitempty"`
	CustomerCode   string         `json:"customer_code,omitempty"`
	Subscriptions  []Subscription `json:"subscriptions,omitempty"`
	Authorizations []interface{}  `json:"authorizations,omitempty"`
	RiskAction     string         `json:"risk_action"`
}

// CustomerList is a list object for customers.
type CustomerList struct {
	Meta   ListMeta
	Values []Customer `json:"data"`
}

// Create creates a new customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
func (s *CustomerService) Create(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customer")
	cust := &Customer{}
	err := s.client.Call("POST", u, customer, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-customer
func (s *CustomerService) Update(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("customer/%d", customer.ID)
	cust := &Customer{}
	err := s.client.Call("PUT", u, customer, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://paystack.com/docs/api/#customer-fetch
func (s *CustomerService) Get(customerCode string) (*Customer, error) {
	u := fmt.Sprintf("/customer/%s", customerCode)
	cust := &Customer{}
	err := s.client.Call("GET", u, nil, cust)

	return cust, err
}

// List returns a list of customers.
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *CustomerService) List() (*CustomerList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of customers
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *CustomerService) ListN(count, offset int) (*CustomerList, error) {
	u := paginateURL("/customer", count, offset)
	cust := &CustomerList{}
	err := s.client.Call("GET", u, nil, cust)
	return cust, err
}

// SetRiskAction can be used to either whitelist or blacklist a customer
// For more details see https://developers.paystack.co/v1.0/reference#whiteblacklist-customer
func (s *CustomerService) SetRiskAction(customerCode, riskAction string) (*Customer, error) {
	reqBody := struct {
		Customer    string `json:"customer"`
		Risk_action string `json:"risk_action"`
	}{
		Customer:    customerCode,
		Risk_action: riskAction,
	}
	cust := &Customer{}
	err := s.client.Call("POST", "/customer/set_risk_action", reqBody, cust)

	return cust, err
}

// DeactivateAuthorization deactivates an authorization
// For more details see https://developers.paystack.co/v1.0/reference#deactivate-authorization
func (s *CustomerService) DeactivateAuthorization(authorizationCode string) (*Response, error) {
	params := url.Values{}
	params.Add("authorization_code", authorizationCode)

	resp := &Response{}
	err := s.client.Call("POST", "/customer/deactivate_authorization", params, resp)

	return resp, err
}
