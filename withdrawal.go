package main

import (
	"strconv"
	"strings"
	"time"
)

var (
	queryIdentifiers = map[string]string{"gte": ">=", "gt": ">", "lte": "<=", "lt": "<", "eq": "="}
	paramDelimiter   = "|"
)

type Withdrawal struct {
	Query     map[string][]string `sql:"-" json:",omitempty"`
	ID        int                 `json:"id"`
	ProductId int                 `json:"product_id"`
	Quantity  int                 `json:"quantity"`
	IssuedAt  int                 `json:"issued_at"`
}

func NewWithdrawl(prod Product, quantity int) *Withdrawal {
	return &Withdrawal{ProductId: prod.ID, Quantity: quantity, IssuedAt: int(time.Now().Unix())}
}

func (w *Withdrawal) Save() error {
	if err := db.Create(w).Error; err != nil {
		return err
	}
	return nil
}

func (w *Withdrawal) Retreive() ([]Withdrawal, error) {
	var ws []Withdrawal
	err := db.Where(*w).Find(&ws, w.BuildQuery()).Error
	return ws, err
}

// /api/query?gte=quantity|200&eq=product_id|4...
// Only uses as connector 'and'
func (w *Withdrawal) BuildQuery() string {
	var query string

	// identifierKey = gte, identifierValue = >=
	for identifierKey, identifierValue := range queryIdentifiers {
		// val = [quantity|200]
		if queryWithKeyValues, ok := w.Query[identifierKey]; ok {

			// currQueryValue  = quantity|200
			for _, currQueryValue := range queryWithKeyValues {
				// splitted = [quantity 200]
				splitted := strings.Split(currQueryValue, paramDelimiter)
				if len(splitted) != 2 {
					continue
				}

				if query != "" {
					query += " and "
				}

				if isNumber(splitted[1]) {
					query += splitted[0] + identifierValue + splitted[1]
				} else {
					query += splitted[0] + identifierValue + "'" + splitted[1] + "'"
				}
			}
		}
	}

	return query
}

func isNumber(s string) bool {
	_, err1 := strconv.Atoi(s)
	_, err2 := strconv.ParseFloat(s, 64)
	if err1 == nil && err2 == nil {
		return true
	}
	return false
}
