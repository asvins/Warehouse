package main

import (
	"errors"
	"fmt"
	"time"
)

type Purchase struct {
	ID             int     `json:"id"`
	CreatedAt      int     `json:"created_at"`
	ConfirmedAt    int     `json:"confirmed_at"`
	ConcludedAt    int     `json:"concluded_at"`
	TotalValue     float64 `json:"total_value"`
	PurschaseOrder Order   `json:"order"`
	OrderId        int     `json:"order_id"`
}

// NewPurchaseFromOrder return a pointer to a newly created struct that uses an order as parameter
func NewPurchaseFromOrder(o *Order) *Purchase {
	totalValue := 0.0
	for _, pp := range o.Pproducts {
		totalValue += pp.Value
	}
	return &Purchase{CreatedAt: int(time.Now().Unix()), TotalValue: totalValue, OrderId: o.ID}
}

// Retreive purchase from database
func (purch *Purchase) Retreive() ([]Purchase, error) {
	var purchs []Purchase
	err := db.Where(purch).Find(&purchs).Error
	if err != nil {
		return nil, err
	}

	for i, p := range purchs {
		order := Order{}
		order.ID = p.OrderId

		if err := db.Model(&p).Related(&order, "OrderId").Error; err != nil {
			return nil, err
		}

		pproducts := []PurchaseProduct{}
		if err := db.Model(order).Related(&pproducts, "Pproducts").Error; err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return nil, err
		}
		order.Pproducts = pproducts

		p.PurschaseOrder = order
		purchs[i] = p
	}

	return purchs, err
}

//Save new purchase into database
func (purch *Purchase) Save() error {
	return db.Create(purch).Error
}

// Update purchase on database
func (purch *Purchase) Update() error {
	return db.Save(purch).Error
}

// Delete purchase on database
func (purch *Purchase) Delete() error {
	return db.Where(purch).Delete(Purchase{}).Error
}

// Confirm purchase
func (purch *Purchase) Confirm() error {
	return db.Model(purch).UpdateColumn(Purchase{ConfirmedAt: int(time.Now().Unix())}).Error
}

// Conclude purchase
func (purch *Purchase) Conclude() error {
	rowsAffected := db.Model(purch).Where("confirmed_at != ?", 0).UpdateColumn(Purchase{ConcludedAt: int(time.Now().Unix())}).RowsAffected
	if rowsAffected != 1 {
		return errors.New("[ERROR] Purchase must be confirmed before conlcuded")
	}

	return nil
}

func (purch *Purchase) RetreiveOpen() ([]Purchase, error) {
	return purch.retreivePlainQuery("confirmed_at = 0 and concluded_at = 0")
}

func (purch *Purchase) RetreiveConfirmed() ([]Purchase, error) {
	return purch.retreivePlainQuery("confirmed_at != 0 and concluded_at = 0")
}

func (purch *Purchase) RetreiveConcluded() ([]Purchase, error) {
	return purch.retreivePlainQuery("confirmed_at != 0 and concluded_at != 0")
}

func (purch *Purchase) retreivePlainQuery(query string) ([]Purchase, error) {
	purchs := []Purchase{}
	err := db.Where(query).Find(&purchs).Error
	return purchs, err
}
