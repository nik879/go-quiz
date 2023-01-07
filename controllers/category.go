package controllers

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gubesch/go-quiz/models"
	"log"
	"net/http"
	"strconv"
)

func ShowAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	categories, err := models.ShowCategories()
	if err != nil {
		NewResponse(false, "SQL error").JSON(w, http.StatusOK)
	} else {
		//_=json.NewEncoder(w).Encode(categories)
		NewResponse(true, "request ok").Attr("categories", categories).JSON(w, http.StatusOK)
	}
}
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	err := category.CreateNewCategory()
	if err != nil {
		NewResponse(false, "SQL error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "successfully created new category").JSON(w, http.StatusOK)
		//_= json.NewEncoder(w).Encode("successfully created new category")
	}
}
func EditCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Wrong Parameters").JSON(w, http.StatusOK)
	} else {
		category.ID = id
		err = category.EditCategory()
		if err != nil {
			NewResponse(false, "SQL error").JSON(w, http.StatusOK)
		} else {
			NewResponse(true, "successfully edited the category").JSON(w, http.StatusOK)
		}
	}

}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	var category models.Category
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Wrong Parameters").JSON(w, http.StatusOK)
	} else {
		category.ID = id
		err = category.DeleteCategory()
		if err != nil {
			NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
		} else {
			NewResponse(true, "successfully deleted the category").JSON(w, http.StatusOK)
		}
	}
}
