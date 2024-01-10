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
	CreateUser()
}

type useDatabase struct {
	databaseConnection *sql.DB
}

func (ur *useDatabase) CreateUser() {
	// Implemente a lógica para criar um usuário no banco de dados aqui
}

/*

package repository

import "go.mongodb.org/mongo-driver/mongo"

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type UserRepository interface {
	CreateUser()
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func (ur *userRepository) CreateUser() {
	// Implemente a lógica para criar um usuário no banco de dados aqui
}



*/
