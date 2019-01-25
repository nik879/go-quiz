package main

import (
	"github.com/gubesch/go-quiz/migration"
	"github.com/joho/godotenv"
	"log"
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
		}
	} else {
		port := os.Getenv("HTTP_PORT")
		address := os.Getenv("LISTEN_ADDR")
		log.Printf("Starting server on %s:%s", address, port)
	}
}
