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

	if len(orders) == 0 {
		return errors.NotFound("record not found")
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

	if len(orders) == 0 {
		return errors.NotFound("record not found")
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

func approveOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	var order Order

	params := r.URL.Query()
	oId, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	order.ID = oId
	if err := order.Approve(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order.ID)
	return nil
}

func cancelOrder(w http.ResponseWriter, r *http.Request) errors.Http {
	var order Order

	params := r.URL.Query()
	oId, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	order.ID = oId
	if err := order.Cancel(); err != nil {
		return errors.BadRequest(err.Error())
	}

	rend.JSON(w, http.StatusOK, order.ID)
	return nil
}
