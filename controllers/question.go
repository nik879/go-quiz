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

func ShowAllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	questions, err := models.GetAllQuestions()
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "Successfully got all questions.").Attr("questions", questions).JSON(w, http.StatusOK)
	}
}
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	var question models.Question
	_ = json.NewDecoder(r.Body).Decode(&question)
	err := question.CreateNewQuestion()
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "Successfully created new question.").JSON(w, http.StatusOK)
	}
}
func EditQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp

	parameters := mux.Vars(r)
	log.Println(parameters)
	var question models.Question
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusOK)
	} else {
		_ = json.NewDecoder(r.Body).Decode(&question)
		question.ID = id
		err = question.EditQuestion()
		if err != nil {
			NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
		} else {
			NewResponse(true, "Successfully edited this question.").JSON(w, http.StatusOK)
		}
	}

}
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	var question models.Question
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusOK)
	} else {
		question.ID = id
		err = question.DeleteQuestion()
		if err != nil {
			NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
		} else {
			NewResponse(true, "Successfully deleted this question.").JSON(w, http.StatusOK)
		}
	}

}
func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	questionID, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusOK)
	}
	answerID, err := strconv.Atoi(parameters["answer_id"])
	if err != nil {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusOK)
	}
	success, err := models.AnswerQuestion(questionID, answerID)
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
	} else {
		NewResponse(true, "Question answer method successfully.").Attr("correct_answer", success).JSON(w, http.StatusOK)
	}
}
func GetRandomQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	//check for timestamp
	parameters := mux.Vars(r)
	log.Println(parameters)
	categoryID, err := strconv.Atoi(parameters["cat_id"])
	if err != nil {
		NewResponse(false, "Wrong parameters").JSON(w, http.StatusOK)
		return
	}
	question, err := models.GetRandomQuestionPerCategory(categoryID)
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
		return
	}
	NewResponse(true, "Successfully got a random question.").Attr("question", question).JSON(w, http.StatusOK)

}

func GetSingleQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoded := context.Get(r, "decoded")
	log.Println(decoded)
	parameters := mux.Vars(r)
	log.Println(parameters)
	questionID, err := strconv.Atoi(parameters["id"])
	if err != nil {
		NewResponse(false, "Question ID must be an int").JSON(w, http.StatusOK)
		return
	}
	question, err := models.GetSpecificQuestion(questionID)
	if err != nil {
		NewResponse(false, "SQL Error").JSON(w, http.StatusOK)
		return
	}
	NewResponse(true, "Successfully found the question.").Attr("question", question).JSON(w, http.StatusOK)
}
