package paystack

import (
	"fmt"
	"net/url"
)

// SubscriptionService handles operations related to the subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type SubscriptionService service

// Subscription represents a Paystack subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type Subscription struct {
	ID          int    `json:"id,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Integration int    `json:"integration,omitempty"`
	// inconsistent API response. Create returns Customer code, Fetch returns an object
	Customer  interface{} `json:"customer,omitempty"`
	Plan      string      `json:"plan,omitempty"`
	StartDate string      `json:"start,omitempty"`
	// inconsistent API response. Fetch returns string, List returns an object
	Authorization    interface{}   `json:"authorization,omitempty"`
	Invoices         []interface{} `json:"invoices,omitempty"`
	Status           string        `json:"status,omitempty"`
	Quantity         int           `json:"quantity,omitempty"`
	Amount           int           `json:"amount,omitempty"`
	SubscriptionCode string        `json:"subscription_code,omitempty"`
	EmailToken       string        `json:"email_token,omitempty"`
	EasyCronID       string        `json:"easy_cron_id,omitempty"`
	CronExpression   string        `json:"cron_expression,omitempty"`
	NextPaymentDate  string        `json:"next_payment_date,omitempty"`
	OpenInvoice      string        `json:"open_invoice,omitempty"`
}

// SubscriptionRequest represents a Paystack subscription request
type SubscriptionRequest struct {
	// customer code or email address
	Customer string `json:"customer,omitempty"`
	// plan code
	Plan          string `json:"plan,omitempty"`
	Authorization string `json:"authorization,omitempty"`
	StartDate     string `json:"start,omitempty"`
}

// SubscriptionList is a list object for subscriptions.
type SubscriptionList struct {
	Meta   ListMeta
	Values []Subscription `json:"data"`
}

// Create creates a new subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
func (s *SubscriptionService) Create(subscription *SubscriptionRequest) (*Subscription, error) {
	u := fmt.Sprintf("/subscription")
	sub := &Subscription{}
	err := s.client.Call("POST", u, subscription, sub)
	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-subscription
func (s *SubscriptionService) Update(subscription *Subscription) (*Subscription, error) {
	u := fmt.Sprintf("subscription/%d", subscription.ID)
	sub := &Subscription{}
	err := s.client.Call("PUT", u, subscription, sub)
	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-subscription
func (s *SubscriptionService) Get(id int) (*Subscription, error) {
	u := fmt.Sprintf("/subscription/%d", id)
	sub := &Subscription{}
	err := s.client.Call("GET", u, nil, sub)
	return sub, err
}

// List returns a list of subscriptions.
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *SubscriptionService) List() (*SubscriptionList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of subscriptions
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *SubscriptionService) ListN(count, offset int) (*SubscriptionList, error) {
	u := paginateURL("/subscription", count, offset)
	sub := &SubscriptionList{}
	err := s.client.Call("GET", u, nil, sub)
	return sub, err
}

// Enable enables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#enable-subscription
func (s *SubscriptionService) Enable(subscriptionCode, emailToken string) (Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := Response{}
	err := s.client.Call("POST", "/subscription/enable", params, &resp)
	return resp, err
}

// Disable disables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#disable-subscription
func (s *SubscriptionService) Disable(subscriptionCode, emailToken string) (Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := Response{}
	err := s.client.Call("POST", "/subscription/disable", params, &resp)
	return resp, err
}
