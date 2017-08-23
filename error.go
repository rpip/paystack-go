package paystack

import (
	"encoding/json"
	"net/url"
)

type Error struct {
	Message        string                 `json:"message,omitempty"`
	HTTPStatusCode int                    `json:"code,omitempty"`
	Details        map[string]interface{} `json:"errors,omitempty"`
	URL            *url.URL               `json:"url,omitempty"`
}

func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
