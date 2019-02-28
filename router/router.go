package router

import (
	"github.com/gorilla/mux"
	"github.com/gubesch/go-quiz/controllers"
	"github.com/gubesch/go-quiz/middleware"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", controllers.LoginUser).Methods("POST"); //done
	router.HandleFunc("/auth/register",controllers.RegisterUser).Methods("POST"); //done

	subrouter := router.PathPrefix("/api").Subrouter()

	categoryRouter := subrouter.PathPrefix("/category").Subrouter()
	categoryRouter.HandleFunc("", controllers.ShowAllCategories).Methods("GET")
	categoryRouter.HandleFunc("/create", controllers.CreateCategory).Methods("POST")
	categoryRouter.HandleFunc("/{id}/edit", controllers.EditCategory).Methods("PUT")
	categoryRouter.HandleFunc("/{id}/delete", controllers.DeleteCategory).Methods("DELETE")

	questionRouter := subrouter.PathPrefix("/question").Subrouter()
	questionRouter.HandleFunc("", controllers.ShowAllQuestions).Methods("GET")
	questionRouter.HandleFunc("/create", controllers.CreateQuestion).Methods("POST")
	questionRouter.HandleFunc("/{id}/edit", controllers.EditQuestion).Methods("PUT")
	questionRouter.HandleFunc("/{id}/delete", controllers.DeleteQuestion).Methods("DELETE")
	questionRouter.HandleFunc("/category/{cat_id}",controllers.GetRandomQuestion).Methods("GET")
	questionRouter.HandleFunc("/{id}/answer",controllers.AnswerQuestion).Methods("POST")

	subrouter.Use(middleware.ValidateMiddleware)

	return router

}
