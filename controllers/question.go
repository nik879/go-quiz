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

func ShowAllQuestions(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	questions,err := models.GetAllQuestions()
	if err != nil{
		NewResponse(false,"SQL Error").JSON(w,http.StatusOK)
	} else {
		NewResponse(true,"Successfully got all questions.").Attr("questions",questions).JSON(w,http.StatusOK)
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
		NewResponse(false,"SQL Error").JSON(w,http.StatusOK)
	} else {
		NewResponse(true,"Successfully created new question.").JSON(w,http.StatusOK)
	}
}
func EditQuestion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp

	parameters := mux.Vars(r)
	var question models.Question
	id,err := strconv.Atoi(parameters["id"])
	if err != nil{
		NewResponse(false,"Wrong parameters").JSON(w,http.StatusOK)
	} else {
		_ = json.NewDecoder(r.Body).Decode(&question)
		question.ID = id
		err = question.DeleteQuestion()
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		}
		err = question.CreateNewQuestion()
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		} else {
			_= json.NewEncoder(w).Encode("successfully edited this question")
		}
	}

}
func DeleteQuestion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	var question models.Question
	id,err := strconv.Atoi(parameters["id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		question.ID = id
		err = question.DeleteQuestion()
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		} else {
			_= json.NewEncoder(w).Encode("successfully deleted this question")
		}
	}

}
func AnswerQuestion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	questionID,err := strconv.Atoi(parameters["id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	}
	answerID,err := strconv.Atoi(parameters["answer_id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	}
	success,err := models.AnswerQuestion(questionID,answerID)
	if err != nil {
		_=json.NewEncoder(w).Encode(err)
	} else {
		_=json.NewEncoder(w).Encode(success)
	}
}
func GetRandomQuestion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r,"decoded")
	fmt.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	categoryID, err := strconv.Atoi(parameters["cat_id"])
	if err != nil{
		_=json.NewEncoder(w).Encode(err)
	} else {
		question,err := models.GetRandomQuestionPerCategory(categoryID)
		if err != nil{
			_=json.NewEncoder(w).Encode(err)
		} else {
			_=json.NewEncoder(w).Encode(question)
		}
	}
}