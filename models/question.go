package models

import (
	"github.com/gubesch/go-quiz/migration"
)

type Question struct {
	ID			int 		`json:"id,omitempty"`
	Question	string 		`json:"question,omitempty"`
	CategoryID	int 		`json:"category_id,omitempty"`
	Answers		[]Answer 	`json:"answers,omitempty"`
}

type Answer struct {
	ID			int 	`json:"id,omitempty"`
	Answer		string 	`json:"answer,omitempty"`
	Correct		bool 	`json:"correct,omitempty"`
	QuestionID	int 	`json:"question_id,omitempty"`
}

func ShowQuestions() (allquestions []Question, err error) {
	db := migration.GetDbInstance()
	questionsQuery, err := db.Query("SELECT * FROM questions;")
	if err != nil{
		return nil,err
	}
	for questionsQuery.Next() {
		var question Question
		err = questionsQuery.Scan(&question.ID, &question.Question, &question.CategoryID)
		if err != nil{
			return nil,err
		}
		allquestions = append(allquestions, question)
	}
	return
}

func (q *Question) CreateNewQuestion() (err error) {
	db := migration.GetDbInstance()
	tx,err := db.Begin()
	InsertQuestionExec,err := tx.Exec("INSERT INTO `questions` (`question`,`categoryID`) VALUES (?,?);", q.Question,q.CategoryID)
	if err != nil {
		tx.Rollback()
		return
	}
	id,err :=InsertQuestionExec.LastInsertId()
	if err != nil{
		return
	}
	for _, element := range q.Answers {
		_,err =tx.Exec("INSERT INTO `answers` (`answer`,`correct`,`questionID`) VALUES (?,?,?);",element.Answer,element.Correct,id)
		if err != nil{
			tx.Rollback()
			return
		}
	}
	err = tx.Commit()
	if err != nil{
		tx.Rollback()
		return
	}

	return nil

}