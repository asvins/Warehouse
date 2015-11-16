package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/decoder"
)

func FillOrderIdWithUrlValue(o *Order, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	o.ID = id
	return nil
}

func retreiveOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	queryString := r.URL.Query()
	var o Order
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeURLValues(&o, queryString); err != nil {
		return errors.BadRequest(err.Error())
	}

	orders, err := o.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(orders) == 0 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, orders)
	return nil
}

func retreiveOrderById(w http.ResponseWriter, r *http.Request) errors.Http {
	o := Order{}

	if err := FillOrderIdWithUrlValue(&o, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	orders, err := o.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(orders) != 1 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, orders[0])
	return nil
}

func retreiveOpenOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	order, err := GetOpenOrder()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order)
	return nil
}

func approveOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	order := Order{}

	if err := FillOrderIdWithUrlValue(&order, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := order.Approve(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order.ID)
	return nil
}

func cancelOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	order := Order{}

	if err := FillOrderIdWithUrlValue(&order, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := order.Cancel(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order.ID)
	return nil
}
