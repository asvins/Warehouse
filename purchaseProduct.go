package main

type PurchaseProduct struct {
	ID        int
	Value     float64
	Quantity  int
	ProductId int
	OrderId   int
}

func NewPurchaseProduct(p *Product) *PurchaseProduct {
	quantity := p.MinQuantity - p.CurrQuantity
	return &PurchaseProduct{Quantity: quantity, ProductId: p.ID}
}
