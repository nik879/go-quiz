package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gubesch/go-quiz/models"
	"net/http"
	"os"
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
			"timestamp": 123123,
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		}
		_=json.NewEncoder(w).Encode(map[string]interface{}{"jwt":tokenString})
	} else {
		_=json.NewEncoder(w).Encode("Login unsuccessful")
	}
}
