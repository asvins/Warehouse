package main

import (
	"net/http"

	"github.com/asvins/utils/responseHelper"
	"github.com/asvins/warehouse/decoder"
)

func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	var p Product

	decoder := decoder.NewDecoder()
	err := decoder.DecodeURLValues(&p, queryString)

	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	products, err := p.Retreive()

	rj := responseHelper.NewResponseJSON(products, err)
	responseHelper.WriteBack(w, r, rj)

}

func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)
	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	err = p.Save()
	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	rj := responseHelper.NewResponseJSON("Product successfully saved", err)
	responseHelper.WriteBack(w, r, rj)
}

func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	err = p.Update()

	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	rj := responseHelper.NewResponseJSON("Product updated successfully", err)
	responseHelper.WriteBack(w, r, rj)
}

func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		rj := responseHelper.NewResponseJSON(nil, err)
		responseHelper.WriteBack(w, r, rj)
		return
	}

	err = p.Delete()

	rj := responseHelper.NewResponseJSON("Product deleted successully", err)
	responseHelper.WriteBack(w, r, rj)
}
