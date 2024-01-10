package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	connectiondb "github.com/eron97/inject.git/connectiondb"
	"github.com/eron97/inject.git/controllers"
	"github.com/eron97/inject.git/repository"
	"github.com/eron97/inject.git/routes"
	service "github.com/eron97/inject.git/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	cfg := connectiondb.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	// Criação de uma conexão MySQL
	db, err := connectiondb.NewConnectionDB(cfg)
	if err != nil {
		fmt.Println("Erro ao conectar ao MySQL:", err)
		return
	}

	defer db.Close()

	userController := initDependencies(db)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	router.Run(":8080")

}

func initDependencies(
	database *sql.DB,
) controllers.ControllerInterface {
	repo := repository.NewDatabase(database)
	service := service.NewDomainService(repo)
	return controllers.NewControllerInterface(service)
}
