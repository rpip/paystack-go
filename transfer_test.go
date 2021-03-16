package paystack

import (
	"testing"
)

func TestInitiateTransfer(t *testing.T) {
	c.Transfer.EnableOTP()

	recipient := &TransferRecipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	recipient1, err := c.Transfer.CreateRecipient(recipient)

	req := &TransferRequest{
		Source:    "balance",
		Reason:    "Delivery pickup",
		Amount:    300,
		Recipient: recipient1.RecipientCode,
	}

	transfer, err := c.Transfer.Initiate(req)

	if err != nil {
		t.Error(err)
	}

	if transfer.TransferCode == "" {
		t.Errorf("Expected transfer code, got %+v", transfer.TransferCode)
	}

	// fetch transfer
	trf, err := c.Transfer.Get(transfer.TransferCode)
	if err != nil {
		t.Error(err)
	}

	if trf.TransferCode == "" {
		t.Errorf("Expected transfer code, got %+v", trf.TransferCode)
	}
}

/* FAILS: Error message: Invalid amount passed
func TestBulkTransfer(t *testing.T) {
	// You need to disable the Transfers OTP requirement to use this endpoint
	c.Transfer.DisableOTP()

	// retrieve the transfer recipient list
	recipients, err := createDemoRecipients()

	if err != nil {
		t.Error(err)
	}

	transfer := &BulkTransfer{
		Source:   "balance",
		Currency: "NGN",
		Transfers: []map[string]interface{}{
			{
				"amount":    50000,
				"recipient": recipients[0].RecipientCode,
			},
			{
				"amount":    50000,
				"recipient": recipients[1].RecipientCode,
			},
		},
	}

	_, err = c.Transfer.MakeBulkTransfer(transfer)

	if err != nil {
		t.Error(err)
	}
}
*/

func TestTransferList(t *testing.T) {
	// retrieve the transfer list
	transfers, err := c.Transfer.List()
	if err != nil {
		t.Errorf("Expected Transfer list, got %d, returned error %v", len(transfers.Values), err)
	}
}

func TestTransferRecipientList(t *testing.T) {
	//fmt.Println("createDemoRecipients <<<<<<<")
	//_, err := createDemoRecipients()

	//if err != nil {
	//	t.Error(err)
	//}

	//fmt.Println("ListRecipients <<<<<<<")
	// retrieve the transfer recipient list
	recipients, err := c.Transfer.ListRecipients()

	if err != nil || !(len(recipients.Values) > 0) || !(recipients.Meta.Total > 0) {
		t.Errorf("Expected Recipients list, got %d, returned error %v", len(recipients.Values), err)
	}
}

func createDemoRecipients() ([]*TransferRecipient, error) {
	recipient1 := &TransferRecipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Carpenter"},
	}

	recipient2 := &TransferRecipient{
		Type:          "Nuban",
		Name:          "Customer 2",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Chef"},
	}

	recipient3 := &TransferRecipient{
		Type:          "Nuban",
		Name:          "Customer 2",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	_, err := c.Transfer.CreateRecipient(recipient1)
	_, err = c.Transfer.CreateRecipient(recipient2)
	_, err = c.Transfer.CreateRecipient(recipient3)

	return []*TransferRecipient{recipient1, recipient2, recipient3}, err
}
