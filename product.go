package main

import (
	"errors"
	"fmt"
)

//Product struct that defines a product
type Product struct {
	ID           int
	Name         string `sql:"size:255"`
	Type         int    // rethink
	Description  string `sql:"size:255"`
	CurrQuantity int
	MinQuantity  int
	CurrValue    float64
}

//Save new product on database
func (p *Product) Save() error {
	if err := db.Create(p).Error; err != nil {
		return err
	}

	if p.NeedRefill() {
		//return AddProductToOpenOrder(*p)
		fmt.Println("[INFO] Adding product to order")
	}

	return nil
}

// Update product on database
func (p *Product) Update() error {
	if err := db.Save(p).Error; err != nil {
		return err
	}

	if p.NeedRefill() {
		// AddProductToOpenOrder(*p)
		fmt.Println("[INFO] Adding product to order")
	} else {
		fmt.Println("[INFO] UPDATE - product doesn't need refill")
		//	has, err := OpenOrderHasProduct(*p)
		//	if err != nil {
		//		return err
		//	}
		//	if has {
		//		RemoveProductFromOpenOrder(*p)
		//	}
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

// NeedRefill verify if product need refill
func (p *Product) NeedRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
