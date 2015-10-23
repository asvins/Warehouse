package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// ResponseJSON ...
type ResponseJSON struct {
	ResponseBody interface{} `json:",omitempty"`
	Error        string      `json:",omitempty"`
}

func writeBack(w http.ResponseWriter, r *http.Request, i interface{}) {
	ct := r.Header.Get("Content-Type")
	switch ct {
	case "application/json":
		bJSON, err := json.Marshal(i)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bJSON)
		break

	case "application/xml":
		bXML, err := xml.Marshal(i)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bXML)
		break

	}
}

// NewResponseJSON ...
func NewResponseJSON(object interface{}, err error) ResponseJSON {
	if err != nil {
		return ResponseJSON{Error: err.Error()}
	}
	return ResponseJSON{ResponseBody: object}
}
