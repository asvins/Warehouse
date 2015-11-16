package main

import "time"

type Withdrawal struct {
	ID        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
	IssuedAt  int `json:"issued_at"`
}

func NewWithdrawl(prod Product, quantity int) *Withdrawal {
	return &Withdrawal{ProductId: prod.ID, Quantity: quantity, IssuedAt: int(time.Now().Unix())}
}

func (w *Withdrawal) Save() error {
	if err := db.Create(w).Error; err != nil {
		return err
	}
	return nil
}

func (w *Withdrawal) Retreive() ([]Withdrawal, error) {
	var ws []Withdrawal
	err := db.Where(*w).Find(&ws).Error
	return ws, err
}
