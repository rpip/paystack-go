[![Build Status](https://travis-ci.org/rpip/paystack-go.svg?branch=master)](https://travis-ci.org/rpip/paystack-go)

# Go library for the Paystack API.

paystack-go is a Go client library for accessing the Paystack API.


## Usage

``` go
import "https://github.com/rpip/paystack-go"

apiKey = "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"
client = paystack.NewClient(apiKey)

cust := &Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+23400000000000000",
}
// create the customer
customer, err := c.Customer.Create(cust)
if err != nil {
    // do something with error
}

// Get customer by ID
client.Customers.Get(customer.ID)

// retrieve list of plans
ch, err := client.Plan.List()
```

See the test files for more examples.

## TODO
- [ ] Documentation
- [ ] More test cases
- [ ] Better handling of API call errors
- [ ] Support request context
- [ ] Test on App Engine

## CONTRIBUTING
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. Please make sure the build succeeds and the test suite passes.
