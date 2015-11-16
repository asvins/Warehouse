package main

import "time"

type Withdrawal struct {
	ID       int     `json:"id"`
	Prod     Product `json:"product"`
	Quantity int     `json:"quantity"`
	IssuedAt int     `json:"issued_at"`
}

func NewWithdrawl(prod Product, quantity int) *Withdrawal {
	return &Withdrawal{Prod: prod, Quantity: quantity, IssuedAt: int(time.Now().Unix())}
}

func (w *Withdrawal) Save() error {
	if err := db.Create(w).Error; err != nil {
		return err
	}
	return nil
}
