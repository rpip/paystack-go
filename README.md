[![Build Status](https://travis-ci.org/rpip/paystack-go.svg?branch=master)](https://travis-ci.org/rpip/paystack-go)

# Go library for the Paystack API.

paystack-go is a Go client library for accessing the Paystack API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the Paystack API documentation at https://developers.paystack.co/v1.0/reference.

## Usage

``` go
import "https://github.com/rpip/paystack-go"

apiKey = "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"
// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
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
- [ ] Upload godocs

## CONTRIBUTING
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. Please make sure the build succeeds and the test suite passes.
