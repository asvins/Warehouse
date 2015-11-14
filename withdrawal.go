package main

type Withdrawal struct {
	ID         int
	Prod       Product
	Quantity   int
	IssuedAt   int
	ApprovedAt int
}
