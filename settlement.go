package paystack

// SettlementService handles operations related to the settlement
// For more details see https://developers.paystack.co/v1.0/reference#create-settlement
type SettlementService service

// SettlementList is a list object for settlements.
type SettlementList struct {
	Meta   ListMeta
	Values []Response `json:"data,omitempty"`
}

// List returns a list of settlements.
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *SettlementService) List() (*SettlementList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of settlements
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *SettlementService) ListN(count, offset int) (*SettlementList, error) {
	u := paginateURL("/settlement", count, offset)
	pg := &SettlementList{}
	err := s.client.Call("GET", u, nil, pg)
	return pg, err
}
