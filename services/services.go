package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eron97/inject.git/models"
	database "github.com/eron97/inject.git/repository"
	"github.com/gin-gonic/gin"
)

type DomainService interface {
	CreateUser(c *gin.Context) error
	ReadUser(c *gin.Context) (models.GetUserID, error)
	ReadAllUsers(c *gin.Context) ([]models.GetUser, error)
	DeleteUser(c *gin.Context) error
}

type useDomainService struct {
	userRepository database.Database
}

func NewDomainService(
	db database.Database,
) DomainService {
	return &useDomainService{db}
}

func (uds *useDomainService) CreateUser(c *gin.Context) error {
	var user models.CreateUser

	if err := c.BindJSON(&user); err != nil {
		return fmt.Errorf("Erro ao realizar o binding do JSON: %s", err.Error())
	}

	if err := uds.userRepository.CreateUser([]models.CreateUser{user}); err != nil {
		return fmt.Errorf("Erro ao criar o usuário no repositório: %s", err.Error())
	}

	return nil
}

func (uds *useDomainService) ReadUser(c *gin.Context) (models.GetUserID, error) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		// Retornar erro caso haja problema com a conversão do ID.
		return models.GetUserID{}, fmt.Errorf("invalid ID: %v", err)
	}

	user, err := uds.userRepository.ReadUser(id)
	if err != nil {
		// Retornar erro caso ocorra um problema ao ler o usuário.
		return models.GetUserID{}, fmt.Errorf("failed to read user: %v", err)
	}

	// Retornar o usuário obtido.
	return user, nil
}

func (uds *useDomainService) ReadAllUsers(c *gin.Context) ([]models.GetUser, error) {
	return uds.userRepository.ReadAllUsers()
}

func (uds *useDomainService) DeleteUser(c *gin.Context) error {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return fmt.Errorf("invalid ID: %v", err)
	}

	if err := uds.userRepository.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir usuário do banco de dados"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso"})
	return nil
}
