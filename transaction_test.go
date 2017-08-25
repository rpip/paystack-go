package paystack

import (
	"fmt"
	"testing"
	"time"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TestInitializeTransaction(t *testing.T) {
	txn := &TransactionRequest{
		Email:     "user123@gmail.com",
		Amount:    6000,
		Reference: "Txn-" + fmt.Sprintf("%d", makeTimestamp()),
	}
	resp, err := c.Transaction.Initialize(txn)
	if err != nil {
		t.Error(err)
	}

	if resp["authorization_code"] == "" {
		t.Error("Missing transaction authorization code")
	}

	if resp["access_code"] == "" {
		t.Error("Missing transaction access code")
	}

	if resp["reference"] == "" {
		t.Error("Missing transaction reference")
	}

	txn1, err := c.Transaction.Verify(resp["reference"].(string))

	if err != nil {
		t.Error(err)
	}

	if txn1.Amount != txn.Amount {
		t.Errorf("Expected transaction amount %f, got %+v", txn.Amount, txn1.Amount)
	}

	if txn1.Reference == "" {
		t.Errorf("Missing transaction reference")
	}

	_, err = c.Transaction.Get(txn1.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestTransactionList(t *testing.T) {
	// retrieve the transaction list
	transactions, err := c.Transaction.List()
	if err != nil {
		t.Errorf("Expected Transaction list, got %d, returned error %v", len(transactions.Values), err)
	}
}

func TestTransactionTotals(t *testing.T) {
	_, err := c.Transaction.Totals()
	if err != nil {
		t.Error(err)
	}
}

func TestExportTransaction(t *testing.T) {
	resp, err := c.Transaction.Export(nil)
	if err != nil {
		t.Error(err)
	}

	if _, ok := resp["path"]; !ok {
		t.Error("Expected transactiion export path")
	}
}
