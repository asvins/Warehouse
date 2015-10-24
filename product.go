package main

import (
	"errors"

	"github.com/asvins/common_db/postgres"
)

const (
	//FOOD ...
	FOOD = 1 << iota
	//CLEANING ...
	CLEANING
	//ROOMITENS ...
	ROOMITENS // towels, bed sheets
)

// ProductToSend ...
type ProductToSend struct {
	ProductID  int     `json:"produto_id"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}

// ProductToConsume ...
type ProductToConsume struct {
	ID       int
	Quantity int
}

//Product struct that defines a product
type Product struct {
	BaseModel    `sql:"-" json:",omitempty"` // Ignore this field
	ID           int
	Name         string `sql:"size:255"`
	Type         int
	Description  string `sql:"size:255"`
	CurrQuantity int
	MinQuantity  int
}

//Save ..
func (p *Product) Save() error {
	db := postgres.GetDatabase()

	err := db.Create(p).Error
	if err != nil {
		return err
	}

	if p.NeedRefill() {
		err = AddProductToOpenOrder(*p)
	}

	return err
}

// Update ...
func (p *Product) Update() error {
	db := postgres.GetDatabase()

	err := db.Save(p).Error
	if err != nil {
		return err
	}

	if p.NeedRefill() {
		AddProductToOpenOrder(*p)
	} else if has, err := OpenOrderHasProduct(*p); has && err == nil {
		RemoveProductFromOpenOrder(*p)
	}

	return err
}

// Delete ...
func (p *Product) Delete() error {
	db := postgres.GetDatabase()
	return db.Delete(p).Error
}

//Retreive ... it uses the object and a plain query to execute sql cmds
func (p *Product) Retreive() ([]Product, error) {
	db := postgres.GetDatabase()
	var query string
	if p.QueryParams != nil {
		query = buildQuery(p.QueryParams)
	}

	orderBy := p.QueryParams["order_by"]

	var products []Product
	var err error
	//Remove queryParams
	p.QueryParams = nil
	if orderBy != "" {
		err = db.Order(orderBy).Where(*p).Find(&products, query).Error
	} else {
		err = db.Where(*p).Find(&products, query).Error
	}

	return products, err
}

// Consume ...
func (p *Product) Consume(quantity int) error {
	db := postgres.GetDatabase()

	var pp Product
	err := db.Where(*p).First(&pp).Error

	if err != nil {
		return err
	}

	if pp.CurrQuantity-quantity < 0 {
		return errors.New("Requested quantity exceeds the available amount")
	}

	pp.CurrQuantity = pp.CurrQuantity - quantity
	err = pp.Update()

	if err != nil {
		return err
	}

	return nil
}

//NeedRefill verify if product need refill
func (p *Product) NeedRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
