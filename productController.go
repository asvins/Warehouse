package main

import (
	"net/http"
	"strconv"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/decoder"
)

func retreiveProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	queryString := r.URL.Query()
	var p Product
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeURLValues(&p, queryString); err != nil {
		return errors.BadRequest(err.Error())
	}

	products, err := p.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(products) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, products)

	return nil
}

func retreiveProductById(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	p := Product{}

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	p.ID = id

	products, err := p.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(products) == 0 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, products)
	return nil
}

func insertProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	var p Product
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Save(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func updateProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	var p Product
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	p.ID = id

	if err := p.Update(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func deleteProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	p := Product{}

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	p.ID = id

	if err := p.Delete(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func consumeProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	p := Product{}

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	p.ID = id

	qt, err := strconv.Atoi(params.Get("quantity"))

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Consume(qt); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}
