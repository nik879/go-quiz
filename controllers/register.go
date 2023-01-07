package controllers

import (
	"encoding/json"
	"github.com/gubesch/go-quiz/models"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Username != "" && user.Password != "" {
		if len(user.Username) >= 3 && len(user.Password) >= 8 {
			err := user.Register()
			if err != nil {
				NewResponse(false, "Username already exists").JSON(w, http.StatusOK)
			} else {
				NewResponse(true, "New User Created").JSON(w, http.StatusCreated)
			}
		} else {
			NewResponse(false, "Username must be at least 3 characters long and password 8!").JSON(w, http.StatusOK)
		}
	} else {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusBadRequest)
	}
}
