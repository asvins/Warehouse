package main

type Purchase struct {
	ID             int     `json:"id"`
	CreatedAt      int     `json:"created_at"`
	ConfirmedAt    int     `json:"confirmed_at"`
	ConcludedAt    int     `json:"concluded_at"`
	TotalValue     float64 `json:"total_value"`
	PurschaseOrder Order   `json:"order"`
}
