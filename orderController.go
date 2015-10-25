package main

import (
	"net/http"

	"github.com/asvins/utils/responseHelper"
	"github.com/asvins/warehouse/decoder"
)

func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o Order
	GetOpenOrder(&o)
	rj := responseHelper.NewResponseJSON(o, nil)
	responseHelper.WriteBack(w, r, rj)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	err = order.Update()
	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	rj := responseHelper.NewResponseJSON("Order updated successfully", err)
	responseHelper.WriteBack(w, r, rj)
}
