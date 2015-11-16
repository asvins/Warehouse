package main

type BaseModel struct {
	QueryParams map[string]string `sql:"-" json:",omitempty"`
}

var (
	queryIdentifiers    = map[string]string{"gte": ">=", "gt": ">", "lte": "<=", "lt": "<", "eq": "="}
	identifierSeparator = "__"
)

/*
*	A idéia é que na query string sejam passados parametros como:
*	/api/query?gte=issued_at__192728471&gte=value__137.2&lte=quantity__20
*
*	O identifierSeparator é que precisa ser pensado direito.
 */
func BuildQueryFromParams(queryMap map[string]string) string {
	//TODO
	return ""
}
