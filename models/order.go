package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//Order is the struct that defines the purchase order
type Order struct {
	ID        int               `json:"id"`
	Pproducts []PurchaseProduct `json:"purchase_products"`
	Approved  bool              `json:"approved"`
	Canceled  bool              `json:"canceled"`
	CreatedAt int               `json:"created_at"`
	ClosedAt  int               `json:"closed_at"`
}

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// ORDER METHODS //////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

// Retreive order from database
func (order *Order) Retreive(db *gorm.DB) ([]Order, error) {
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
func (order *Order) Save(db *gorm.DB) error {
	return db.Create(order).Error
}

// Update order on database
func (order *Order) Update(db *gorm.DB) error {
	return db.Save(order).Error
}

func (order *Order) Approve(db *gorm.DB) error {
	if err := db.Model(order).UpdateColumn(Order{Approved: true, ClosedAt: int(time.Now().Unix())}).Error; err != nil {
		return err
	}

	orders, err := order.Retreive(db)
	if err != nil {
		return err
	}

	if len(orders) != 1 {
		return errors.New("[ERROR] Query for recently approved order failed")
	}

	NewPurchaseFromOrder(&orders[0]).Save(db)
	return nil
}

func (order *Order) Cancel(db *gorm.DB) error {
	return db.Model(order).UpdateColumn(Order{Canceled: true, ClosedAt: int(time.Now().Unix())}).Error
}

// Delete order from database
func (order *Order) Delete(db *gorm.DB) error {
	return db.Delete(order).Error
}

// HasProduct verify if the given order has the specific product
func (order *Order) HasProduct(db *gorm.DB, product Product) (bool, error) {
	if err := db.Model(order).Association("Products").Find(&product).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (order *Order) AddProduct(db *gorm.DB, pproduct *PurchaseProduct) error {
	queryObj := PurchaseProduct{ProductId: pproduct.ProductId, OrderId: order.ID}

	pps, err := queryObj.Retreive(db)
	if err != nil {
		return err
	}

	pproduct.OrderId = order.ID
	if len(pps) == 0 {
		return db.Model(order).Association("Pproducts").Append([]PurchaseProduct{*pproduct}).Error
	} else if len(pps) == 1 {
		pproduct.ID = pps[0].ID
		return db.Save(pproduct).Error
	} else {
		return errors.New("[ERROR] Fatal Error.. database is corrupted")
	}

}

// RemoveProduct removes a product from the order
func (order *Order) RemoveProduct(db *gorm.DB, pproduct PurchaseProduct) error {
	return db.Where(&pproduct).Delete(&pproduct).Error
}

// createAndAddProduct will create a new order an insert the given product in it
func (order *Order) createAndAddProduct(db *gorm.DB, pproduct *PurchaseProduct) error {
	fmt.Println("[DEBUG] WILL CREATE NEW ORDER BEFORE INSERTING")
	order.CreatedAt = int(time.Now().Unix())
	if err := db.Create(order).Error; err != nil {
		return err
	}

	return order.AddProduct(db, pproduct)
}

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// ORDER FUNCTIONS /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

//GetOpenOrder returns an open order if there is one on database
func GetOpenOrder(db *gorm.DB) (*Order, error) {
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
func AddProductToOpenOrder(db *gorm.DB, pproduct *PurchaseProduct) error {
	order, err := GetOpenOrder(db)
	if err != nil {
		if err.Error() == "record not found" {
			order = &Order{}
			return order.createAndAddProduct(db, pproduct)
		}
		return err
	}
	return order.AddProduct(db, pproduct)
}

// the order must have a PurchaseProduct of type Product ...
func OpenOrderHasProduct(db *gorm.DB, pproduct PurchaseProduct) (*Order, error) {
	order, err := GetOpenOrder(db)
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
