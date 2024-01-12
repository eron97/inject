package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eron97/inject.git/models"
)

func NewDatabase(database *sql.DB) Database {
	return &useDatabase{
		database,
	}
}

type Database interface {
	CreateUser(users []models.CreateUser) error
	CreateVerificEmail(email string) (bool, error)
	ReadUser(userID int) (models.GetUserID, error)
	ReadAllUsers() ([]models.GetUser, error)
	DeleteUser(userID int) error
}

type useDatabase struct {
	databaseConnection *sql.DB
}

func (con *useDatabase) CreateVerificEmail(email string) (bool, error) {
	var count int

	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	row := con.databaseConnection.QueryRow(query, email)

	if err := row.Scan(&count); err != nil {
		// Ocorreu um erro ao escanear o resultado da query
		return false, fmt.Errorf("erro ao verificar email: %v", err)
	}

	if count != 0 {
		return true, nil
	}

	return false, nil
}

func (con *useDatabase) CreateUser(users []models.CreateUser) error {

	for _, user := range users {
		_, err := con.databaseConnection.Exec(
			"INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
			user.First_Name, user.Last_Name, user.Email, user.Password,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (con *useDatabase) ReadUser(userID int) (models.GetUserID, error) {
	query := "SELECT username, first_name, last_name FROM users WHERE id = ?"

	row := con.databaseConnection.QueryRow(query, userID)

	var user models.GetUserID

	err := row.Scan(&user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		log.Println(err)
		return models.GetUserID{}, err
	}

	return user, nil
}

func (con *useDatabase) ReadAllUsers() ([]models.GetUser, error) {
	rows, err := con.databaseConnection.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []models.GetUser

	for rows.Next() {
		var user models.GetUser
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.BirthDate, &user.PhoneNumber)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (con *useDatabase) DeleteUser(userID int) error {
	_, err := con.databaseConnection.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
