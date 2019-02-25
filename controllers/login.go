package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gubesch/go-quiz/models"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	loginSuccessful,err := user.Login()

	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	}

	if loginSuccessful {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
		})
		tokenString, err := token.SignedString([]byte("go-quiz"))
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		}
		_=json.NewEncoder(w).Encode(tokenString)
	} else {
		_=json.NewEncoder(w).Encode("Login unsuccessful")
	}
}
