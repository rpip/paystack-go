package paystack

import (
	"testing"
)

func TestChargeServiceCreate(t *testing.T) {
	bankAccount := BankAccount{
		Code:          "057",
		AccountNumber: "0000000000",
	}

	charge := ChargeRequest{
		Email:    "your_own_email_here@gmail.com",
		Amount:   10000,
		Bank:     &bankAccount,
		Birthday: "1999-12-31",
	}

	resp, err := c.Charge.Create(&charge)
	if err != nil {
		t.Errorf("Create Charge returned error: %v", err)
	}

	if resp["reference"] == "" {
		t.Error("Missing transaction reference")
	}
}

func TestChargeServiceCheckPending(t *testing.T) {
	bankAccount := BankAccount{
		Code:          "057",
		AccountNumber: "0000000000",
	}

	charge := ChargeRequest{
		Email:    "your_own_email_here@gmail.com",
		Amount:   10000,
		Bank:     &bankAccount,
		Birthday: "1999-12-31",
	}

	resp, err := c.Charge.Create(&charge)
	if err != nil {
		t.Errorf("Create charge returned error: %v", err)
	}

	if resp["reference"] == "" {
		t.Error("Missing charge reference")
	}

	resp2, err := c.Charge.CheckPending(resp["reference"].(string))
	if err != nil {
		t.Errorf("Check pending charge returned error: %v", err)
	}

	if resp2["status"] == "" {
		t.Error("Missing charge pending status")
	}

	if resp2["reference"] == "" {
		t.Error("Missing charge pending reference")
	}
}
