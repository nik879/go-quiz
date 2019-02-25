package models

import (
	"github.com/gubesch/go-quiz/migration"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type User struct {
	ID			int		`json:"id,omitempty"`
	Username	string 	`json:"username,omitempty"`
	Password	string 	`json:"password,omitempty"`
}

func (u *User) Register() (err error) {
	db := migration.NewDbConnection()
	stmtRegister,err := db.Prepare("INSERT INTO `users` (`username`, `pw_hash`) VALUES (?,?)")
	if err != nil {
		return err
	}

	passwordHash,err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	_,err = stmtRegister.Exec(u.Username,passwordHash)
	if err != nil {
		return err
	}

	err = stmtRegister.Close()
	if err != nil {
		return err
	}

	return nil;
}

func (u *User) Login() (b bool,err error){
	db := migration.NewDbConnection()
	userQuery,err := db.Query("SELECT username, pw_hash FROM users WHERE username = ?",u.Username)
	if err != nil{
		return false,err
	}
	for userQuery.Next(){
		var user User
		err = userQuery.Scan(&user.Username,&user.Password)
		if err != nil{
			return false,err
		}
		validUser := CheckPasswordHash(u.Password,user.Password)
		return validUser,nil
	}
	return false,nil
}