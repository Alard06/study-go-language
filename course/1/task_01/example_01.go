package main

import (
	"customer"
)

func example01() {
	var firstUser customer.Customer
	firstUser.Name = "Alex"
	firstUser.Debt = 3000

	secontUser := customer.Customer{
		Name: "Bob",
		Debt: 2000,
	}

}
