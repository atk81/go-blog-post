package model

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/atk81/go-blog-post/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func GetAllUser() ([]User, error) {
	db := utils.GetDB()
	sqlQuery := "SELECT * FROM users"
	var Users []User
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		Users = append(Users, user)
	}
	return Users, err

}

func RegisterNewUser(username, password string) error {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return errors.New("username and password cannot be empty")
	} else if UserExists(username) {
		return errors.New("username already exists")
	}
	fmt.Println("Registering new user with username: ", username)
	db := utils.GetDB()
	var id uuid.UUID
	sqlQuery := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(sqlQuery, username, password).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	return err
}

func UserExists(username string) bool {
	db := utils.GetDB()
	sqlQuery := "SELECT * FROM users WHERE username = $1"
	var user User
	row := db.QueryRow(sqlQuery, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err == nil {
		return true
	}
	return false
}

func IsUserValid(username, password string) bool {
	db := utils.GetDB()
	sqlQuery := "SELECT * FROM users WHERE username = $1 AND password = $2"
	var user User
	row := db.QueryRow(sqlQuery, username, password)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err == nil {
		return true
	}
	return false
}
