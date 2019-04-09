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
	categoryRouter.HandleFunc("", controllers.ShowAllCategories).Methods("GET") //done
	categoryRouter.HandleFunc("/create", controllers.CreateCategory).Methods("POST") //done
	categoryRouter.HandleFunc("/{id}/edit", controllers.EditCategory).Methods("PUT") //done
	categoryRouter.HandleFunc("/{id}/delete", controllers.DeleteCategory).Methods("DELETE") //done

	questionRouter := subrouter.PathPrefix("/question").Subrouter()
	questionRouter.HandleFunc("", controllers.ShowAllQuestions).Methods("GET") //done
	questionRouter.HandleFunc("/create", controllers.CreateQuestion).Methods("POST") //done
	questionRouter.HandleFunc("/{id}/edit", controllers.EditQuestion).Methods("PUT") //done
	questionRouter.HandleFunc("/{id}/delete", controllers.DeleteQuestion).Methods("DELETE") //done
	questionRouter.HandleFunc("/category/{cat_id}",controllers.GetRandomQuestion).Methods("GET") //done
	questionRouter.HandleFunc("/{id}/answer/{answer_id}",controllers.AnswerQuestion).Methods("GET")

	subrouter.Use(middleware.ValidateMiddleware)

	return router

}
