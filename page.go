package paystack

import "fmt"

// PageService handles operations related to the page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
type PageService service

// Page represents a Paystack page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
type Page struct {
	ID           int                 `json:"id,omitempty"`
	CreatedAt    string              `json:"createdAt,omitempty"`
	UpdatedAt    string              `json:"updatedAt,omitempty"`
	Domain       string              `json:"domain,omitempty"`
	Integration  int                 `json:"integration,omitempty"`
	Name         string              `json:"name,omitempty"`
	Slug         string              `json:"slug,omitempty"`
	Description  string              `json:"description,omitempty"`
	Amount       float32             `json:"amount,omitempty"`
	Currency     string              `json:"currency,omitempty"`
	Active       bool                `json:"active,omitempty"`
	RedirectURL  string              `json:"redirect_url,omitempty"`
	CustomFields []map[string]string `json:"custom_fields,omitempty"`
}

// PageList is a list object for pages.
type PageList struct {
	Meta   ListMeta
	Values []Page `json:"data,omitempty"`
}

// Create creates a new page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
func (s *PageService) Create(page *Page) (*Page, error) {
	u := fmt.Sprintf("/page")
	pg := &Page{}
	err := s.client.Call("POST", u, page, pg)

	return pg, err
}

// Update updates a page's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-page
func (s *PageService) Update(page *Page) (*Page, error) {
	u := fmt.Sprintf("page/%d", page.ID)
	pg := &Page{}
	err := s.client.Call("PUT", u, page, pg)

	return pg, err
}

// Get returns the details of a page.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-page
func (s *PageService) Get(id int) (*Page, error) {
	u := fmt.Sprintf("/page/%d", id)
	pg := &Page{}
	err := s.client.Call("GET", u, nil, pg)

	return pg, err
}

// List returns a list of pages.
// For more details see https://developers.paystack.co/v1.0/reference#list-pages
func (s *PageService) List() (*PageList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of pages
// For more details see https://developers.paystack.co/v1.0/reference#list-pages
func (s *PageService) ListN(count, offset int) (*PageList, error) {
	u := paginateURL("/page", count, offset)
	pg := &PageList{}
	err := s.client.Call("GET", u, nil, pg)
	return pg, err
}
