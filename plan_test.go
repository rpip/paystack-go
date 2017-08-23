package paystack

import "testing"

func TestPlanCRUD(t *testing.T) {
	plan1 := &Plan{
		Name:     "Monthly retainer",
		Interval: "monthly",
		Amount:   500000,
	}

	// create the plan
	plan, err := c.Plan.Create(plan1)
	if err != nil {
		t.Errorf("CREATE Plan returned error: %v", err)
	}

	if plan.PlanCode == "" {
		t.Errorf("Expected Plan code to be set")
	}

	// retrieve the plan
	plan, err = c.Plan.Get(plan.ID)
	if err != nil {
		t.Errorf("GET Plan returned error: %v", err)
	}

	if plan.Name != plan1.Name {
		t.Errorf("Expected Plan Name %v, got %v", plan.Name, plan1.Name)
	}

	// retrieve the plan list
	plans, err := c.Plan.List()
	if err != nil || !(len(plans.Values) > 0) || !(plans.Meta.Total > 0) {
		t.Errorf("Expected Plan list, got %d, returned error %v", len(plans.Values), err)
	}
}
