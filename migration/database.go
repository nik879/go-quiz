package migration

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func NewDbConnection() sql.DB {

	//function for creating a database connection
	driver := os.Getenv("DB_CONNECTION")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	domain := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")

	database, err := sql.Open(driver, username + ":" + password + "@" + domain + "/" + dbName)
	if err != nil {
		log.Fatal(err)
	}
	return *database
}

func DropDatabase(){
	path,_ := filepath.Abs("migration/dropDB.sql")
	file,err:= ioutil.ReadFile(path)
	if err != nil{
		fmt.Println(err)
	}
	executeSQLFile(string(file))
}

func MigrateDatabase(){
	path,_ := filepath.Abs("migration/quiz.sql")
	file,err:= ioutil.ReadFile(path)
	if err != nil{
		fmt.Println(err)
	}
	executeSQLFile(string(file))
}

func executeSQLFile(fileString string){
	statements := strings.Split(fileString, ";")
	db := NewDbConnection()
	defer db.Close()
	for _, stmt := range statements {
		if strings.Contains(stmt, "EXISTS"){
			_,err := db.Exec(string(stmt) + ";")
			if err != nil{
				fmt.Println(err)
			}
		}
	}
}
