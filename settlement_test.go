package paystack

import (
	"fmt"
	"testing"
)

func TestSettlementList(t *testing.T) {
	// retrieve the settlement list
	settlements, err := c.Settlement.List()

	if err != nil {
		t.Error(err)
	}

	if err == nil {
		fmt.Printf("Settlements total: %d", len(settlements.Values))
	}
}
