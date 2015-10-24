package main

import (
	"net/http"

	"github.com/asvins/utils/responseHelper"
)

func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o Order
	GetOpenOrder(&o)
	rj := responseHelper.NewResponseJSON(o, nil)
	responseHelper.WriteBack(w, r, rj)
}
