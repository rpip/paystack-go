package paystack

import (
	"testing"
)

func TestChargeService_Create(t *testing.T) {
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

	resp, error := c.Charge.Create(&charge)
	if error != nil {
		t.Error(error)
	}

	if resp["reference"] == "" {
		t.Error("Missing transaction reference")
	}
}

func TestChargeService_CheckPending(t *testing.T) {
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

	resp, error := c.Charge.Create(&charge)
	if error != nil {
		t.Error(error)
	}

	if resp["reference"] == "" {
		t.Error("Missing charge reference")
	}

	resp2, error := c.Charge.CheckPending(resp["reference"].(string))
	if error != nil {
		t.Error(error)
	}

	if resp2["status"] == "" {
		t.Error("Missing charge pending status")
	}

	if resp2["reference"] == "" {
		t.Error("Missing charge pending reference")
	}
}