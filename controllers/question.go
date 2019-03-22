package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gubesch/go-quiz/models"
	"net/http"
)

func ShowAllQuestions(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	questions,err := models.ShowQuestions()
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		_=json.NewEncoder(w).Encode(questions)
	}
}
func CreateQuestion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	var question models.Question
	_ = json.NewDecoder(r.Body).Decode(&question)
	err := question.CreateNewQuestion()
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		_= json.NewEncoder(w).Encode("successfully created new question")
	}
}
func EditQuestion(w http.ResponseWriter, r *http.Request){

}
func DeleteQuestion(w http.ResponseWriter, r *http.Request){

}
func AnswerQuestion(w http.ResponseWriter, r *http.Request){

}
func GetRandomQuestion(w http.ResponseWriter, r *http.Request){

}