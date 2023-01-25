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
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) Register() (err error) {
	db := migration.GetDbInstance()
	//defer db.Close()
	stmtRegister, err := db.Prepare("INSERT INTO `users` (`username`, `pw_hash`) VALUES (?,?);")
	defer stmtRegister.Close()
	if err != nil {
		return err
	}

	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmtRegister.Exec(u.Username, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() (allUsers []User, err error) {
	db := migration.GetDbInstance()
	userQuery, err := db.Query("SELECT username FROM users;")
	if err != nil {
		return nil, err
	}
	for userQuery.Next() {
		var user User
		err = userQuery.Scan(&user.Username)
		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, user)
	}
	return
}

func (u *User) Login() (b bool, err error) {
	db := migration.GetDbInstance()
	//defer db.Close()
	userQuery, err := db.Query("SELECT username, pw_hash FROM users WHERE username = ?;", u.Username)
	defer userQuery.Close()
	if err != nil {
		return false, err
	}
	for userQuery.Next() {
		var user User
		err = userQuery.Scan(&user.Username, &user.Password)
		if err != nil {
			return false, err
		}
		validUser := CheckPasswordHash(u.Password, user.Password)
		return validUser, nil
	}
	return false, nil
}

func (u *User) DeleteUser() (err error) {

	db := migration.GetDbInstance()
	//defer db.Close()
	stmtDeleteCategory, err := db.Prepare("DELETE FROM users WHERE username=?;")
	defer stmtDeleteCategory.Close()
	if err != nil {
		return
	}
	_, err = stmtDeleteCategory.Exec(u.Username)
	if err != nil {
		return
	}
	return nil
}
