package main

import (
	"net/http"

	"github.com/asvins/router/errors"
	"github.com/asvins/utils/responseHelper"
	"github.com/asvins/warehouse/decoder"
)

func GETProductHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	queryString := r.URL.Query()
	var p Product

	decoder := decoder.NewDecoder()
	err := decoder.DecodeURLValues(&p, queryString)

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	products, err := p.Retreive()

	rj := responseHelper.NewResponseJSON(products, err)
	responseHelper.WriteBack(w, r, rj)
	return nil
}

func POSTProductHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	err = p.Save()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rj := responseHelper.NewResponseJSON("Product successfully saved", err)
	responseHelper.WriteBack(w, r, rj)
	return nil
}

func PUTProductHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	err = p.Update()

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rj := responseHelper.NewResponseJSON("Product updated successfully", err)
	responseHelper.WriteBack(w, r, rj)
	return nil
}

func DELETEProductHandler(w http.ResponseWriter, r *http.Request) errors.Http {
	var p Product
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	err = p.Delete()

	rj := responseHelper.NewResponseJSON("Product deleted successully", err)
	responseHelper.WriteBack(w, r, rj)
	return nil
}
