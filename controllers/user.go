package controllers

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gubesch/go-quiz/models"
	"log"
	"net/http"
)

func ShowAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	users, err := models.GetAllUsers()
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "Successfully got all users.").Attr("users", users).JSON(w, http.StatusOK)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	var user models.User
	username := parameters["username"]

	user.Username = username
	err := user.DeleteUser()
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "successfully deleted the user").JSON(w, http.StatusOK)
	}
}
