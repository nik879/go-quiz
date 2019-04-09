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

func GetAllQuestions() (allquestions []Question, err error) {
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

func AnswerQuestion(questionID int, answerID int) (success bool, err error){
	db:= migration.GetDbInstance()
	answerQuestionQuery := db.QueryRow("SELECT correct FROM answers WHERE questionID=? AND ID=?;",questionID, answerID)
	err = answerQuestionQuery.Scan(&success)
	return
}

func GetRandomQuestionPerCategory(categoryID int) (randomQuestion Question, err error){
	db := migration.GetDbInstance()
	randQuestionQuery, err := db.Query("SELECT * FROM questions WHERE categoryID=? ORDER BY RAND() LIMIT 1;",categoryID)
	defer randQuestionQuery.Close()
	if err != nil{
		return
	}
	randQuestionQuery.Next()
	err = randQuestionQuery.Scan(&randomQuestion.ID,&randomQuestion.Question,&randomQuestion.CategoryID)
	if err != nil{
		return
	}
	answersQuery,err := db.Query("SELECT ID,answer FROM answers WHERE questionID=?;",randomQuestion.ID)
	if err != nil{
		return
	}
	for answersQuery.Next() {
		var answer Answer
		err = answersQuery.Scan(&answer.ID,&answer.Answer)
		if err != nil{
			return
		}
		randomQuestion.Answers = append(randomQuestion.Answers, answer)
	}
	return
}


func (q *Question) DeleteQuestion() (err error){
	db := migration.GetDbInstance()
	_, err = db.Exec("DELETE FROM `questions` WHERE ID=?;",q.ID)
	if err != nil{
		return
	}
	return nil
}

func (q *Question) CreateNewQuestion() (err error) {
	db := migration.GetDbInstance()
	tx,err := db.Begin()
	InsertQuestionExec,err := tx.Exec("INSERT INTO `questions` (`question`,`categoryID`) VALUES (?,?);", q.Question,q.CategoryID)
	if err != nil {
		_=tx.Rollback()
		return
	}
	id,err :=InsertQuestionExec.LastInsertId()
	if err != nil{
		return
	}
	for _, element := range q.Answers {
		_,err =tx.Exec("INSERT INTO `answers` (`answer`,`correct`,`questionID`) VALUES (?,?,?);",element.Answer,element.Correct,id)
		if err != nil{
			_=tx.Rollback()
			return
		}
	}
	err = tx.Commit()
	if err != nil{
		_=tx.Rollback()
		return
	}

	return nil

}