package main

import (
	"net/http"

	"github.com/asvins/router/errors"
	"github.com/asvins/warehouse/models"
)

func retreiveWithdrawal(w http.ResponseWriter, r *http.Request) errors.Http {
	withd := models.Withdrawal{}
	if err := BuildStructFromQueryString(&withd, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	withd.Query = r.URL.Query()

	ws, err := withd.Retreive(db)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(ws) == 0 {
		return errors.NotFound("record not found")
	}

	rend.JSON(w, http.StatusOK, ws)

	return nil
}
