package main

import (
	"fmt"
	"time"
)

//Order is the struct that defines the purchase order
type Order struct {
	ID        int               `json:"id"`
	Pproducts []PurchaseProduct `json:"purchase_products"`
	Approved  bool              `json:"approved"`
	CreatedAt int               `json:"created_at"`
	ClosedAt  int               `json:"closed_at"`
}

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// ORDER METHODS //////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

// Retreive order from database
func (order *Order) Retreive() ([]Order, error) {
	var orders []Order
	err := db.Where(*order).Find(&orders).Error

	for i, o := range orders {
		pproducts := []PurchaseProduct{}
		if err := db.Model(o).Related(&pproducts, "Pproducts").Error; err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return nil, err
		}
		o.Pproducts = pproducts
		orders[i] = o
	}

	return orders, err
}

//Save order on database
func (order *Order) Save() error {
	return db.Create(order).Error
}

// Update order on database
func (order *Order) Update() error {
	return db.Save(order).Error
}

func (order *Order) Approve() error {
	return db.Model(order).Debug().UpdateColumn("approved", "true").Error
}

// Delete order from database
func (order *Order) Delete() error {
	err := db.Delete(order).Error
	return err
}

// HasProduct verify if the given order has the specific product
func (order *Order) HasProduct(product Product) (bool, error) {
	if err := db.Model(order).Association("Products").Find(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (order *Order) AddProduct(pproduct *PurchaseProduct) error {
	pproduct.OrderId = order.ID
	return db.Model(order).Association("Pproducts").Append([]PurchaseProduct{*pproduct}).Error
}

// RemoveProduct removes a product from the order
func (order *Order) RemoveProduct(pproduct PurchaseProduct) error {
	return db.Where(&pproduct).Delete(&pproduct).Error
}

// createAndAddProduct will create a new order an insert the given product in it
func (order *Order) createAndAddProduct(pproduct *PurchaseProduct) error {
	order.CreatedAt = int(time.Now().Unix())
	if err := db.Create(order).Error; err != nil {
		return err
	}

	return order.AddProduct(pproduct)
}

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// ORDER FUNCTIONS /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

//GetOpenOrder returns an open order if there is one on database
func GetOpenOrder() (*Order, error) {
	order := Order{}
	if err := db.Where("approved = ?", false).First(&order).Error; err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return nil, err
	}

	pproducts := []PurchaseProduct{}
	if err := db.Model(order).Related(&pproducts, "Pproducts").Error; err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return nil, err
	}
	order.Pproducts = pproducts

	return &order, nil
}

// AddProduct to the existing open order or creates a new order if it needs
func AddProductToOpenOrder(pproduct *PurchaseProduct) error {
	order, err := GetOpenOrder()
	if err != nil {
		if err.Error() == "record not found" {
			order = &Order{}
			return order.createAndAddProduct(pproduct)
		}
		return err
	}
	return order.AddProduct(pproduct)
}

// the order must have a PurchaseProduct of type Product ...
func OpenOrderHasProduct(pproduct PurchaseProduct) (*Order, error) {
	order, err := GetOpenOrder()
	if err != nil {
		return nil, err
	}

	if err = db.Model(order).Association("Pproducts").Find(&pproduct).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return order, nil
}
