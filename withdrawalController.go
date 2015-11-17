package main

import (
	"net/http"

	"github.com/asvins/router/errors"
)

func retreiveWithdrawal(w http.ResponseWriter, r *http.Request) errors.Http {
	withd := Withdrawal{}
	if err := BuildStructFromQueryString(&withd, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	withd.Query = r.URL.Query()

	ws, err := withd.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(ws) == 0 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, ws)

	return nil
}
