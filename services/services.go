package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eron97/inject.git/models"
	database "github.com/eron97/inject.git/repository"
	"github.com/eron97/inject.git/services/password"
	"github.com/gin-gonic/gin"
)

type DomainService interface {
	CreateUser(request models.CreateUser) string
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

func (service *useDomainService) CreateUser(request models.CreateUser) string {

	emailExists, err := service.userRepository.VerificEmailExist(request.Email)
	if err != nil {
		return "Erro ao verificar a existência do e-mail"
	}

	if !emailExists {
		newPassword, err := password.HashPassword(request.Password)
		if err != nil {
			return "Erro ao criptografar a senha"
		}

		request.Password = newPassword

		err = service.userRepository.CreateUser([]models.CreateUser{request})
		if err == nil {
			return "Usuário criado com sucesso!"
		} else {
			return "Erro ao criar usuário no banco de dados"
		}
	}

	return "E-mail já existe e está associado a outra conta"
}

/*
func (service *useDomainService) CreateUser(request models.CreateUser) string {
	if emailExists, err := service.userRepository.VerificEmailExist(request.Email); err == nil && !emailExists {
		if newPassword, err := password.HashPassword(request.Password); err == nil {
			request.Email = newPassword
			return "Senha criptografada com sucesso!"
		}
	}

	err := service.userRepository.CreateUser([]models.CreateUser{request})
	if err == nil {
		return "Usuário criado com sucesso!"
	} else {
		return "Erro ao criar usuário no banco de dados"
	}
}
*/

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
