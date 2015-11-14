package main

import (
	"net/http"
	"strconv"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/decoder"
)

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

	rend.JSON(w, http.StatusOK, orders)
	return nil
}

func retreiveOrderById(w http.ResponseWriter, r *http.Request) errors.Http {
	params := r.URL.Query()
	o := Order{}

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	o.ID = id

	orders, err := o.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, orders)
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

func updateOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	var order Order
	decoder := decoder.NewDecoder()

	if err := decoder.DecodeReqBody(&order, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	params := r.URL.Query()
	oId, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	order.ID = oId

	if err := order.Update(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order)
	return nil
}
