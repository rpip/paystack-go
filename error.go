package paystack

import (
	"fmt"
	"net/url"
)

type Error struct {
	Message        string                 `json:"message,omitempty"`
	HTTPStatusCode int                    `json:"code,omitempty"`
	Errors         map[string]interface{} `json:"errors,omitempty"`
	URL            *url.URL               `json:"url,omitempty"`
}

func (r *Error) Error() string {
	return fmt.Sprintf("%d %s: %v %+v",
		r.HTTPStatusCode, r.Message, r.URL, r.Errors)
}
