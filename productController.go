package main

import (
	"net/http"

	"github.com/asvins/router/errors"
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

	rend.JSON(w, http.StatusOK, products)
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

	rend.JSON(w, http.StatusOK, "Product successfully saved")
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

	rend.JSON(w, http.StatusOK, "Product updated successfully")
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

	rend.JSON(w, http.StatusOK, "Product deleted successfully")
	return nil
}
