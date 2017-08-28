package paystack

import "testing"

func TestSubscriptionCRUD(t *testing.T) {
	cust := &Customer{
		FirstName: "User123",
		LastName:  "AdminUser",
		Email:     "user123-subscription@gmail.com",
		Phone:     "+23400000000000000",
	}
	// create the customer
	customer, err := c.Customer.Create(cust)
	if err != nil {
		t.Errorf("CREATE Subscription Customer returned error: %v", err)
	}

	plan1 := &Plan{
		Name:     "Monthly subscription retainer",
		Interval: "monthly",
		Amount:   250000,
	}

	// create the plan
	plan, err := c.Plan.Create(plan1)
	if err != nil {
		t.Errorf("CREATE Plan returned error: %v", err)
	}

	subscription1 := &SubscriptionRequest{
		Customer: customer.CustomerCode,
		Plan:     plan.PlanCode,
	}

	// create the subscription
	_, err = c.Subscription.Create(subscription1)
	if err == nil {
		t.Errorf("Expected CREATE Subscription to fail with aunthorized customer, got %+v", err)
	}

	// retrieve the subscription list
	subscriptions, err := c.Subscription.List()
	if err != nil {
		t.Errorf("Expected Subscription list, got %d, returned error %v", len(subscriptions.Values), err)
	}
}
