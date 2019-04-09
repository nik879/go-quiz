package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response map[string]interface{}

func NewResponse(success bool, message string, codes ...int) Response {
	r := Response{}
	r["success"] = success
	r["message"] = message

	if len(codes) >= 1 {
		r["code"] = codes[0]
	}
	if len(codes) == 2 {
		r["statusCode"] = codes[1]
	}

	return r
}


// JSON marshals the Response struct to a JSON string and sets the HTTP Statuscode
func (r Response) JSON(w http.ResponseWriter, statusCode ...int) {
	if len(statusCode) == 1 {
		ArbitraryJSON(w, r, statusCode[0])
	} else {
		_, ok := r["statusCode"]
		if ok {
			code := r["statusCode"].(int)
			// remove statuscode from response map to not render it
			delete(r, "statusCode")
			ArbitraryJSON(w, r, code)
		} else {
			ArbitraryJSON(w, r, http.StatusOK)
		}
	}
}

func (r Response) Attr(key string, value interface{}) Response {
	r[key] = value
	return r
}

// ArbitraryJSON takes an interface value and a statuscode and writes it to a given http.ResponseWriter
func ArbitraryJSON(w http.ResponseWriter, value interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		log.Println(err)
	}
}