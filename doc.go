/*
Package paystack provides the binding for Paystack REST APIs.
Where possible, the services available on the client groups the API into
logical chunks and correspond to the structure of the Paystack API
documentation at https://developers.paystack.co/v1.0/reference.

Usage:

	import "github.com/rpip/paystack-go"

	apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

	// second param is an optional http client, allowing overriding of the HTTP client to use.
	// This is useful if you're running in a Google AppEngine environment
	// where the http.DefaultClient is not available.
	client := paystack.NewClient(apiKey)

	recipient := &TransferRecipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0100000010",
		BankCode:      "044",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	recipient1, err := client.Transfer.CreateRecipient(recipient)

	req := &TransferRequest{
		Source:    "balance",
		Reason:    "Delivery pickup",
		Amount:    30,
		Recipient: recipient1.RecipientCode,
	}

	transfer, err := client.Transfer.Initiate(req)

	// retrieve list of plans
	plans, err := client.Plan.List()

	for i, plan := range plans.Values {
	  fmt.Printf("%+v", plan)
	}

*/
package paystack
