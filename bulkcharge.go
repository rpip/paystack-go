package paystack

import "fmt"

// BulkChargeService handles operations related to the bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
type BulkChargeService service

// BulkChargeBatch represents a bulk charge batch object
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
type BulkChargeBatch struct {
	ID            int    `json:"id,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
	BatchCode     string `json:"batch_code,omitempty"`
	Status        string `json:"status,omitempty"`
	Integration   int    `json:"integration,omitempty"`
	Domain        string `json:"domain,omitempty"`
	TotalCharges  string `json:"total_charges,omitempty"`
	PendingCharge string `json:"pending_charge,omitempty"`
}

// BulkChargeRequest is an array of objects with authorization codes and amount
type BulkChargeRequest struct {
	Items []BulkItem
}

// BulkItem represents a single bulk charge request item
type BulkItem struct {
	Authorization string  `json:"authorization,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
}

// BulkChargeBatchList is a list object for bulkcharges.
type BulkChargeBatchList struct {
	Meta   ListMeta
	Values []BulkChargeBatch `json:"data,omitempty"`
}

// Initiate initiates a new bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
func (s *BulkChargeService) Initiate(req *BulkChargeRequest) (*BulkChargeBatch, error) {
	bulkcharge := &BulkChargeBatch{}
	err := s.client.Call("POST", "/bulkcharge", req.Items, bulkcharge)
	return bulkcharge, err
}

// List returns a list of bulkcharges.
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *BulkChargeService) List() (*BulkChargeBatchList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of bulkcharges
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *BulkChargeService) ListN(count, offset int) (*BulkChargeBatchList, error) {
	u := paginateURL("/bulkcharge", count, offset)
	bulkcharges := &BulkChargeBatchList{}
	err := s.client.Call("GET", u, nil, bulkcharges)
	return bulkcharges, err
}

// Get returns a bulk charge batch
// This endpoint retrieves a specific batch code.
// It also returns useful information on its progress by way of
// the total_charges and pending_charges attributes.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-bulk-charge-batch
func (s *BulkChargeService) Get(idCode string) (*BulkChargeBatch, error) {
	u := fmt.Sprintf("/bulkcharge/%s", idCode)
	bulkcharge := &BulkChargeBatch{}
	err := s.client.Call("GET", u, nil, bulkcharge)
	return bulkcharge, err
}

// GetBatchCharges returns charges in a batch
// This endpoint retrieves the charges associated with a specified batch code.
// Pagination parameters are available. You can also filter by status.
// Charge statuses can be pending, success or failed.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-charges-in-a-batch
func (s *BulkChargeService) GetBatchCharges(idCode string) (Response, error) {
	u := fmt.Sprintf("/bulkcharge/%s/charges", idCode)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)
	return resp, err
}

// PauseBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#pause-bulk-charge-batch
func (s *BulkChargeService) PauseBulkCharge(batchCode string) (Response, error) {
	u := fmt.Sprintf("/bulkcharge/pause/%s", batchCode)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)

	return resp, err
}

// ResumeBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#resume-bulk-charge-batch
func (s *BulkChargeService) ResumeBulkCharge(batchCode string) (Response, error) {
	u := fmt.Sprintf("/bulkcharge/resume/%s", batchCode)
	resp := Response{}
	err := s.client.Call("GET", u, nil, &resp)

	return resp, err
}
