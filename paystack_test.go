package paystack

import "testing"

var c *Client

func init() {
	apiKey := mustGetTestKey()
	c = NewClient(apiKey, nil)
}

func TestResolveCardBIN(t *testing.T) {
	resp, err := c.ResolveCardBIN(59983)
	if err != nil {
		t.Error(err)
	}
	if _, ok := resp["bin"]; !ok {
		t.Errorf("Expected response to contain bin")
	}
}

func TestCheckBalance(t *testing.T) {
	resp, err := c.CheckBalance()
	if err != nil {
		t.Error(err)
	}
	if _, ok := resp["currency"]; !ok {
		t.Errorf("Expected response to contain currency")
	}

	if _, ok := resp["balance"]; !ok {
		t.Errorf("Expected response to contain balance")
	}
}

func TestSessionTimeout(t *testing.T) {
	resp, err := c.GetSessionTimeout()
	if err != nil {
		t.Error(err)
	}
	if _, ok := resp["payment_session_timeout"]; !ok {
		t.Errorf("Expected response to contain payment_session_timeout")
	}

	/*
		// actual tests in the Paystack API console also fails. Likely a server error
			resp, err = c.UpdateSessionTimeout(30)
			if err != nil {
				t.Error(err)
			}

			if _, ok := resp["payment_session_timeout"]; !ok {
				t.Errorf("Expected response to contain payment_session_timeout")
			}
	*/
}
