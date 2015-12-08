package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/models"
)

func FillPurchaseProductIdWithUrlValue(p *models.PurchaseProduct, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func FillPurchaseProductProductIdWithUrlValue(p *models.PurchaseProduct, params url.Values) error {
	id, err := strconv.Atoi(params.Get("product_id"))
	if err != nil {
		return err
	}
	p.ProductId = id

	return nil
}

func retreivePurchaseProducts(w http.ResponseWriter, r *http.Request) errors.Http {
	pp := models.PurchaseProduct{}
	if err := BuildStructFromQueryString(&pp, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	pproducts, err := pp.Retreive(db)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if len(pproducts) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, pproducts)

	return nil
}

func retreivePurchaseProductsByProductId(w http.ResponseWriter, r *http.Request) errors.Http {
	pp := models.PurchaseProduct{}

	if err := FillPurchaseProductProductIdWithUrlValue(&pp, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	pproducts, err := pp.Retreive(db)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if len(pproducts) != 1 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, pproducts[0])
	return nil
}

func retreivePurchaseProductsById(w http.ResponseWriter, r *http.Request) errors.Http {
	pp := models.PurchaseProduct{}

	if err := FillPurchaseProductIdWithUrlValue(&pp, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	pproducts, err := pp.Retreive(db)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if len(pproducts) != 1 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, pproducts[0])
	return nil
}

func updatePurchaseProductOnQuantity(w http.ResponseWriter, r *http.Request) errors.Http {
	pp := models.PurchaseProduct{}

	if err := FillPurchaseProductIdWithUrlValue(&pp, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		errors.BadRequest(err.Error())
	}

	if err := pp.UpdateQuantity(db, quantity); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, pp)
	return nil
}

func updatePurchaseProductOnValue(w http.ResponseWriter, r *http.Request) errors.Http {
	pp := models.PurchaseProduct{}

	if err := FillPurchaseProductIdWithUrlValue(&pp, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	value, err := strconv.ParseFloat(r.URL.Query().Get("value"), 64)
	if err != nil {
		errors.BadRequest(err.Error())
	}

	if err := pp.UpdateValue(db, value); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, pp)
	return nil
}
