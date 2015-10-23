package models

import "strings"

//FIXME - USE PROC ENV VARIABLE TO HOLD THIS VALUE
var comprasIP = "http://192.168.1.132:8080"

//BaseModel struct for all models
type BaseModel struct {
	QueryParams map[string]string `sql:"-" json:",omitempty"`
}

//Identifiers for query
//The ideia is to use JSONS like:
// {
// 	last_modified_gt: 423424354
// 	min_quantity_lte: 10
// 	curr_quantity_eq: "min_quantity"
// 	.
//	.
//	.
// }
var queryIdentifiers = map[string]string{"_gte": ">=", "_gt": ">", "_lte": "<=", "_lt": "<", "_eq": "="}

//Aux functions
func buildQuery(queryMap map[string]string) string {
	var query string
	for k, value := range queryMap {
		for keyIdentifier, vIdentifier := range queryIdentifiers {
			if strings.Contains(k, keyIdentifier) {
				if query != "" {
					query += " and "
				}
				columnName := k[0 : len(k)-len(keyIdentifier)]
				query += columnName + vIdentifier + value
				break
			}
		}
	}

	return query
}
