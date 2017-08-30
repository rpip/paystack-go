package paystack

import "fmt"

// PlanService handles operations related to the plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
type PlanService service

// Plan represents a
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
type Plan struct {
	ID                int     `json:"id,omitempty"`
	CreatedAt         string  `json:"createdAt,omitempty"`
	UpdatedAt         string  `json:"updatedAt,omitempty"`
	Domain            string  `json:"domain,omitempty"`
	Integration       int     `json:"integration,omitempty"`
	Name              string  `json:"name,omitempty"`
	Description       string  `json:"description,omitempty"`
	PlanCode          string  `json:"plan_code,omitempty"`
	Amount            float32 `json:"amount,omitempty"`
	Interval          string  `json:"interval,omitempty"`
	SendInvoices      bool    `json:"send_invoices,omitempty"`
	SendSMS           bool    `json:"send_sms,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	InvoiceLimit      float32 `json:"invoice_limit,omitempty"`
	HostedPage        string  `json:"hosted_page,omitempty"`
	HostedPageURL     string  `json:"hosted_page_url,omitempty"`
	HostedPageSummary string  `json:"hosted_page_summary,omitempty"`
}

// PlanList is a list object for Plans.
type PlanList struct {
	Meta   ListMeta
	Values []Plan `json:"data"`
}

// Create creates a new plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
func (s *PlanService) Create(plan *Plan) (*Plan, error) {
	u := fmt.Sprintf("/plan")
	plan2 := &Plan{}
	err := s.client.Call("POST", u, plan, plan2)
	return plan2, err
}

// Update updates a plan's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-plan
func (s *PlanService) Update(plan *Plan) (Response, error) {
	u := fmt.Sprintf("plan/%d", plan.ID)
	resp := Response{}
	err := s.client.Call("PUT", u, plan, &resp)
	return resp, err
}

// Get returns the details of a plan.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-plan
func (s *PlanService) Get(id int) (*Plan, error) {
	u := fmt.Sprintf("/plan/%d", id)
	plan2 := &Plan{}
	err := s.client.Call("GET", u, nil, plan2)
	return plan2, err
}

// List returns a list of plans.
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *PlanService) List() (*PlanList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of plans
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *PlanService) ListN(count, offset int) (*PlanList, error) {
	u := paginateURL("/plan", count, offset)
	plan2 := &PlanList{}
	err := s.client.Call("GET", u, nil, plan2)
	return plan2, err
}
