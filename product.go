package main

import (
	"errors"
	"fmt"
)

//Product struct that defines a product
type Product struct {
	ID            int               `json:"id"`
	Name          string            `json:"name" sql:"size:255"`
	Type          int               `json:"type"`
	Description   string            `json:"description" sql:"size:255"`
	CurrQuantity  int               `json:"curr_quantity"`
	MinQuantity   int               `json:"min_quantity"`
	PurchProducts []PurchaseProduct `json:"purchase_products"`
	Withdrawals   []Withdrawal      `json:"withdrawals"`
}

//Save new product on database
func (p *Product) Save() error {
	if err := db.Create(p).Error; err != nil {
		return err
	}

	if p.CurrQuantity < p.MinQuantity {
		fmt.Println("[INFO] Adding product to order")
		pp := NewPurchaseProduct(p)
		return AddProductToOpenOrder(pp)
	}

	return nil
}

// Update product on database
func (p *Product) Update() error {
	if err := db.Save(p).Error; err != nil {
		return err
	}

	if p.CurrQuantity < p.MinQuantity {
		pp := NewPurchaseProduct(p)
		return AddProductToOpenOrder(pp)
	} else {
		pp := &PurchaseProduct{ProductId: p.ID}
		order, err := OpenOrderHasProduct(*pp)
		if err != nil {
			return err
		}
		if order != nil {
			pp.OrderId = order.ID
			order.RemoveProduct(*pp)
		}
	}
	return nil
}

// Delete product on database
func (p *Product) Delete() error {
	return db.Where(p).Delete(Product{}).Error
}

// Retreive product on database
func (p *Product) Retreive() ([]Product, error) {
	var products []Product
	err := db.Where(*p).Find(&products).Error

	return products, err
}

// Consume product using the id provided and quantity
// if issued quantity > current quantity an error will be returned
func (p *Product) Consume(quantity int) error {
	var pp Product

	if err := db.Where(*p).First(&pp).Error; err != nil {
		return err
	}

	if pp.CurrQuantity-quantity < 0 {
		return errors.New("Requested quantity exceeds the available amount")
	}

	pp.CurrQuantity = pp.CurrQuantity - quantity

	if err := pp.Update(); err != nil {
		return err
	}

	return nil
}
