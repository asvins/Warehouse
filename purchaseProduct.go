package main

type PurchaseProduct struct {
	ID        int
	Value     float64
	Quantity  int
	ProductId int
	OrderId   int
}

func NewPurchaseProduct(p *Product) *PurchaseProduct {
	return &PurchaseProduct{Quantity: p.MinQuantity - p.CurrQuantity, ProductId: p.ID}
}

func (pp *PurchaseProduct) Retreive() ([]PurchaseProduct, error) {
	var pproducts []PurchaseProduct
	return pproducts, db.Where(*pp).Find(&pproducts).Error
}

func (pp *PurchaseProduct) RetreiveOne() error {
	return db.Where(*pp).First(pp).Error
}
