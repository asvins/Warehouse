package main

import (
	"fmt"

	"github.com/asvins/common_db/postgres"
)

// OrderToSend ...
type OrderToSend struct {
	Products []ProductToSend `json:"produtos"`
}

//Order is the struct that defines the purchase order
type Order struct {
	BaseModel `sql:"-" json:",omitempty"` // Ignore this field
	ID        int
	Products  []Product `gorm:"many2many:order_products;"`
	Valor     int       `json:"valor" sql:"-"`
	Approved  bool
}

// GetByID ...
func (order *Order) GetByID(id int) error {
	db := postgres.GetDatabase()
	order.ID = id

	products := []Product{}
	err := db.Model(order).Related(&products, "Products").Error
	order.Products = products

	fmt.Println(order)
	return err
}

//Save ..
func (order *Order) Save() error {
	db := postgres.GetDatabase()
	return db.Create(order).Error
}

// Update ...
func (order *Order) Update() error {
	//TODO
	return nil
}

// Delete ...
func (order *Order) Delete() error {
	db := postgres.GetDatabase()

	err := db.Delete(order).Error
	return err
}

//GetOpenOrder ...
func GetOpenOrder(order *Order) error {
	db := postgres.GetDatabase()
	err := db.Where("approved = ?", false).First(order).Error
	if err != nil {
		return err
	}
	products := []Product{}
	err = db.Model(order).Related(&products, "Products").Error
	order.Products = products

	return err
}

// OpenOrderHasProduct ...
func OpenOrderHasProduct(product Product) (bool, error) {
	db := postgres.GetDatabase()

	order := Order{}
	err := GetOpenOrder(&order)
	if err != nil {
		return false, err
	}

	err = db.Model(order).Association("Products").Find(&product).Error

	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// RemoveProductFromOpenOrder from the existing opened order
func RemoveProductFromOpenOrder(product Product) error {
	db := postgres.GetDatabase()
	order := Order{}
	err := GetOpenOrder(&order)
	if err != nil {
		return err
	}
	return db.Model(order).Association("Products").Delete([]Product{product}).Error
}

// AddProductToOpenOrder to the existing opened order or creates a new order if needed
func AddProductToOpenOrder(product Product) error {
	var order Order
	err := GetOpenOrder(&order)
	if err != nil {
		if err.Error() == "record not found" {
			return order.createOrderAndAddProduct(product)
		}
		return err
	}
	return order.addProduct(product)
}

func (order *Order) createOrderAndAddProduct(product Product) error {
	db := postgres.GetDatabase()

	err := db.Create(order).Error
	if err != nil {
		return err
	}

	err = db.Model(order).Association("Products").Append([]Product{product}).Error
	if err != nil {
		return err
	}

	return nil
}

func (order *Order) addProduct(product Product) error {
	db := postgres.GetDatabase()

	return db.Model(order).Association("Products").Append([]Product{product}).Error
}
