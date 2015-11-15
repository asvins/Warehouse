package main

type Withdrawal struct {
	ID         int     `json:"id"`
	Prod       Product `json:"product"`
	Quantity   int     `json:"quantity"`
	IssuedAt   int     `json:"issued_at"`
	ApprovedAt int     `json:"approved_at"`
}
