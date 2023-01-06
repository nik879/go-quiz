package main

import (
	"github.com/gubesch/go-quiz/migration"
	"github.com/gubesch/go-quiz/router"
	"github.com/joho/godotenv"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file present. Environment not loaded from file")
	}

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 0 {
		if argsWithoutProg[0] == "--migrate"{
			migration.MigrateDatabase()
		} else if argsWithoutProg[0] == "--migrate:rollback"{
			migration.DropDatabase()
		} else if argsWithoutProg[0] == "--migrate:fresh"{
			migration.DropDatabase()
			migration.MigrateDatabase()
		}
	} else {
		port := os.Getenv("HTTP_PORT")
		address := os.Getenv("LISTEN_ADDR")
		log.Printf("Starting server on %s:%s", address, port)
		log.Fatal(http.ListenAndServe(address + ":" + port, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"}))(router.CreateRouter())))
	}
}
