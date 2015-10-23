package inventory

import (
	"net/http"

	"github.com/asvins/warehouse/models"
	"github.com/asvins/warehouse/responseHelper"
)

func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	models.GetOpenOrder(&o)
	rj := responseHelper.NewResponseJSON(o, nil)
	responseHelper.WriteBack(w, r, rj)
}
