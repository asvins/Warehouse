package main

import "errors"

type PurchaseProduct struct {
	ID        int     `json:"id"`
	Value     float64 `json:"value"`
	Quantity  int     `json:"quantity"`
	ProductId int     `json:"product_id"`
	OrderId   int     `json:"order_id"`
}

func NewPurchaseProduct(p *Product) *PurchaseProduct {
	return &PurchaseProduct{Quantity: p.MinQuantity - p.CurrQuantity, ProductId: p.ID}
}

func VerifyUpdatePurchaseProduct(pp *PurchaseProduct) error {
	p := &Purchase{OrderId: pp.OrderId}
	ps, err := p.Retreive()
	if err != nil {
		return err
	}

	if len(ps) == 0 {
		return nil
	}

	if len(ps) != 1 {
		return errors.New("[ERROR] Error while verifying purchase product update attempt")
	}

	if ps[0].ConfirmedAt != 0 {
		return errors.New("[ERROR] Can't change a Purchase Product of a purchase already confirmed")
	}
	return nil
}

func (pp *PurchaseProduct) Save() error {
	return db.Create(pp).Error
}

func (pp *PurchaseProduct) Update() error {
	return db.Save(pp).Error
}

func (pp *PurchaseProduct) Retreive() ([]PurchaseProduct, error) {
	var pproducts []PurchaseProduct
	return pproducts, db.Where(*pp).Find(&pproducts).Error
}

func (pp *PurchaseProduct) UpdateQuantity(quantity int) error {
	pps, err := pp.Retreive()
	if err != nil {
		return err
	}

	if len(pps) != 1 {
		return errors.New("record not found")
	}

	*pp = pps[0]

	if err := VerifyUpdatePurchaseProduct(pp); err != nil {
		return err
	}

	err = db.Model(pp).UpdateColumn(PurchaseProduct{Quantity: quantity}).Error
	if err != nil {
		return err
	}

	pp.Quantity = quantity

	return nil
}

func (pp *PurchaseProduct) UpdateValue(value float64) error {
	pps, err := pp.Retreive()
	if err != nil {
		return err
	}

	if len(pps) != 1 {
		return errors.New("record not found")
	}

	*pp = pps[0]

	if err := VerifyUpdatePurchaseProduct(pp); err != nil {
		return err
	}

	err = db.Model(pp).UpdateColumn(PurchaseProduct{Value: value}).Error
	if err != nil {
		return err
	}

	pp.Value = value

	return nil
}
