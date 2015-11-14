package main

type Purchase struct {
	ID             int
	CreatedAt      int
	ConfirmedAt    int
	ConcludedAt    int
	TotalValue     float64
	PurschaseOrder Order
}
