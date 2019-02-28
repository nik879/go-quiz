package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gubesch/go-quiz/models"
	"net/http"
	"strconv"
)

func ShowAllCategories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	categories,err := models.ShowCategories()
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		_=json.NewEncoder(w).Encode(categories)
	}
}
func CreateCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	err := category.CreateNewCategory()
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		_= json.NewEncoder(w).Encode("successfully created new category")
	}
}
func EditCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	var category models.Category
	_=json.NewDecoder(r.Body).Decode(&category)
	id,err := strconv.Atoi(parameters["id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		category.ID = id
		err = category.EditCategory()
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		} else {
			_= json.NewEncoder(w).Encode("successfully edited the category")
		}
	}

}
func DeleteCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	var category models.Category
	id,err := strconv.Atoi(parameters["id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		category.ID = id
		err = category.DeleteCategory()
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		} else {
			_= json.NewEncoder(w).Encode("successfully deleted the category")
		}
	}
}