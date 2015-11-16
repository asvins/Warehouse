package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
)

func FillProductIdWithUrlValue(p *Product, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func retreiveProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p := Product{}
	if err := BuildStructFromQueryString(&p, r.URL.Query()); err != nil {
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
	p := Product{}

	if err := FillProductIdWithUrlValue(&p, r.URL.Query()); err != nil {
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

func insertProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p := Product{}
	if err := BuildStructFromReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Save(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func updateProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p := Product{}

	if err := BuildStructFromReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillProductIdWithUrlValue(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Update(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func deleteProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p := Product{}
	if err := FillProductIdWithUrlValue(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Delete(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func consumeProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p := Product{}
	params := r.URL.Query()
	if err := FillProductIdWithUrlValue(&p, params); err != nil {
		return errors.BadRequest(err.Error())
	}

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
