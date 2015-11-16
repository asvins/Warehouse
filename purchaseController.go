package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
)

func FillPurchaseIdWithUrlValue(p *Purchase, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func retreivePurchase(w http.ResponseWriter, r *http.Request) errors.Http {
	purchase := &Purchase{}

	if err := BuildStructFromQueryString(purchase, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	purchases, err := purchase.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(purchases) == 0 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, purchases)

	return nil

}

func retreivePurchaseById(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	purchase := Purchase{}

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	purchase.ID = id

	purchases, err := purchase.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(purchases) != 1 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, purchases[0])
	return nil
}

func retreivePurchaseByOrderId(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	purchase := Purchase{}

	orderId, err := strconv.Atoi(params.Get("order_id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	purchase.OrderId = orderId

	purchases, err := purchase.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(purchases) != 1 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, purchases[0])
	return nil
}

func retreiveOpenPurchase(w http.ResponseWriter, r *http.Request) errors.Http {
	purchase := Purchase{}
	purchs, err := purchase.RetreiveOpen()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, purchs)
	return nil
}

func retreiveConfirmedPurchases(w http.ResponseWriter, r *http.Request) errors.Http {
	purchase := Purchase{}
	purchs, err := purchase.RetreiveConfirmed()

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, purchs)
	return nil
}

func retreiveConcludedPurchases(w http.ResponseWriter, r *http.Request) errors.Http {
	purchase := Purchase{}
	purchs, err := purchase.RetreiveConcluded()

	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, purchs)
	return nil
}

func confirmPurchase(w http.ResponseWriter, r *http.Request) errors.Http {
	var purchase Purchase

	if err := FillPurchaseIdWithUrlValue(&purchase, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := purchase.Confirm(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, purchase.ID)
	return nil
}

func concludePurchase(w http.ResponseWriter, r *http.Request) errors.Http {
	var purchase Purchase

	if err := FillPurchaseIdWithUrlValue(&purchase, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := purchase.Conclude(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, purchase.ID)
	return nil
}
