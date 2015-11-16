package main

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/decoder"
)

func BuildProductFromQueryString(queryString url.Values) (*Product, error) {
	var p Product
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeURLValues(&p, queryString); err != nil {
		return nil, err
	}

	return &p, nil
}

func BuildProductFromReqBody(body io.ReadCloser) (*Product, error) {
	var p Product
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeReqBody(&p, body); err != nil {
		return nil, err
	}
	return &p, nil
}

func FillIdWithUrlValue(p *Product, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func retreiveProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p, err := BuildProductFromQueryString(r.URL.Query())
	if err != nil {
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

	if err := FillIdWithUrlValue(&p, r.URL.Query()); err != nil {
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
	p, err := BuildProductFromReqBody(r.Body)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Save(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}

func updateProduct(w http.ResponseWriter, r *http.Request) errors.Http {
	p, err := BuildProductFromReqBody(r.Body)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillIdWithUrlValue(p, r.URL.Query()); err != nil {
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
	if err := FillIdWithUrlValue(&p, r.URL.Query()); err != nil {
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
	if err := FillIdWithUrlValue(&p, params); err != nil {
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
