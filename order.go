package main

import (
	"fmt"
	"time"
)

//Order is the struct that defines the purchase order
type Order struct {
	ID        int
	Pproducts []PurchaseProduct
	Approved  bool
	CreatedAt int
	ClosedAt  int
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
		products := []Product{}
		if err := db.Model(o).Related(&products, "Products").Error; err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return nil, err
		}
		o.Products = products
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
	if err := db.Save(order).Error; err != nil {
		return err
	}

	return nil
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

func (order *Order) AddProduct(product Product) error {
	return db.Model(order).Association("Products").Append([]Product{product}).Error
}

// RemoveProduct removes a product from the order
func (order *Order) RemoveProduct(product Product) error {
	return db.Model(order).Association("Products").Delete([]Product{product}).Error
}

// createAndAddProduct will create a new order an insert the given product in it
func (order *Order) createAndAddProduct(product Product) error {
	order.CreatedAt = int(time.Now().Unix())
	if err := db.Create(order).Error; err != nil {
		return err
	}

	if err := db.Model(order).Association("Products").Append([]Product{product}).Error; err != nil {
		return err
	}

	return nil
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

	products := []Product{}
	if err := db.Model(order).Related(&products, "Products").Error; err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return nil, err
	}
	order.Products = products

	return &order, nil
}

// AddProduct to the existing open order or creates a new order if it needs
func AddProductToOpenOrder(product Product) error {
	order, err := GetOpenOrder()
	if err != nil {
		if err.Error() == "record not found" {
			order = &Order{}
			return order.createAndAddProduct(product)
		}
		return err
	}
	return order.AddProduct(product)
}

func OpenOrderHasProduct(product Product) (*Order, error) {
	order, err := GetOpenOrder()
	if err != nil {
		return nil, err
	}

	if err = db.Model(order).Association("Products").Find(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return order, nil
}
