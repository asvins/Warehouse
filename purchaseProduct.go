package main

type PurchaseProduct struct {
	ID        int
	value     float64
	quantity  int
	ProductId int
	OrderId   int
}

func NewPurchaseProduct(p *Product) *PurchaseProduct {
	return &PurchaseProduct{quantity: (p.MinQuantity - p.CurrQuantity), ProductId: p.ID}
}
