package main

import (
	"net/http"

	"github.com/asvins/router/errors"
	"github.com/asvins/utils/responseHelper"
	"github.com/asvins/warehouse/decoder"
)

func GETOrderHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	var o Order
	GetOpenOrder(&o)
	rj := responseHelper.NewResponseJSON(o, nil)
	responseHelper.WriteBack(w, r, rj)
	return nil
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	var order Order
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	err = order.Update()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rj := responseHelper.NewResponseJSON("Order updated successfully", err)
	responseHelper.WriteBack(w, r, rj)
	return nil
}
