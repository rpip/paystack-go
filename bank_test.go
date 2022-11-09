package paystack

import "testing"

func TestBankList(t *testing.T) {
	// retrieve the bank list
	banks, err := c.Bank.List()

	if err != nil || !(len(banks.Values) > 0) {
		t.Errorf("Expected Bank list, got %d, returned error %v", len(banks.Values), err)
	}
}

func TestMatchBVN(t *testing.T) {
	req := &BVNRequest{
		AccountNumber: "0001234560",
		BankCode:      "058",
		FirstName:     "Customer",
		LastName:      "2",
		MiddleName:    "1",
	}
	// Test invlaid BVN.
	// Err not nill. Resp status code is 400
	req.BVN = "21212917"
	resp, err := c.Bank.MatchBVN(req)
	if err == nil {
		t.Errorf("Expected error for invalid BVN, got %+v'", resp)
	}

	// Test free calls limit
	// Error is nil
	// &{Meta:{CallsThisMonth:0 FreeCallsLeft:0} BVN:cZ+MKrsLAqJCUi+hxIdQqw==}’
	req.BVN = "21212917741"
	resp, err = c.Bank.MatchBVN(req)
	if resp.Meta.FreeCallsLeft != 0 {
		t.Errorf("Expected free calls limit exceeded, got %+v'", resp)
	}
	// TODO(yao): Reproduce error: Your balance is not enough to fulfill this request
}

func TestResolveAccountNumber(t *testing.T) {
	resp, err := c.Bank.ResolveAccountNumber("0022728151", "063")
	if err == nil {
		t.Errorf("Expected error, got %+v'", resp)
	}

	/*
		if _, ok := resp["account_number"]; !ok {
			t.Errorf("Expected response to contain 'account_number'")
		}

		if _, ok := resp["account_name"]; !ok {
			t.Errorf("Expected response to contain 'account_name'")
		}
	*/
}
