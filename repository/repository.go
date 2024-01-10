package repository

import (
	"database/sql"
)

func NewDatabase(database *sql.DB) Database {
	return &useDatabase{
		database,
	}
}

type Database interface {
	CreateUser() error
	ReadlAllUsers() error
}

type useDatabase struct {
	databaseConnection *sql.DB
}

func (ur *useDatabase) CreateUser() error {
	return nil
}

func (ur *useDatabase) ReadlAllUsers() error {
	return nil
}
