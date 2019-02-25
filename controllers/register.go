package controllers

import (
	"encoding/json"
	"github.com/gubesch/go-quiz/models"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_=json.NewDecoder(r.Body).Decode(&user)
	if user.Username != "" && user.Password != "" {
		if len(user.Username) >= 3 && len(user.Password) >= 8 {
			err := user.Register()
			if err != nil {
				_=json.NewEncoder(w).Encode(err)
			} else {
				_=json.NewEncoder(w).Encode("Created User")
			}
		} else {
			_=json.NewEncoder(w).Encode("error code")
		}
	} else {
		_=json.NewEncoder(w).Encode("wrong params")
	}
}
