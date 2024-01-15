package controllers

import (
	"net/http"

	"github.com/eron97/inject.git/config/validate"
	"github.com/eron97/inject.git/models"
	service "github.com/eron97/inject.git/services"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	ReadlAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	ReadUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type useControllerInterface struct {
	service service.DomainService
}

func NewControllerInterface(
	service service.DomainService,
) ControllerInterface {
	return &useControllerInterface{
		service: service,
	}
}

func (controller *useControllerInterface) CreateUser(c *gin.Context) {
	var request models.CreateUser

	err := validate.ValidateRequest(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro de validação"})
		return
	}

	resp := controller.service.CreateUser(request)
	c.JSON(http.StatusOK, gin.H{"message": resp})
}

func (pkg *useControllerInterface) ReadlAllUsers(c *gin.Context) {

	resp, err := pkg.service.ReadAllUsers(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": resp})
}

func (uci *useControllerInterface) ReadUser(c *gin.Context) {
	resp, err := uci.service.ReadUser(c)
	if err != nil {
		// Trate o erro conforme necessário, por exemplo, enviando uma resposta de erro.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Processar resp conforme necessário.
	c.JSON(http.StatusOK, gin.H{"users": resp})
}

func (uci *useControllerInterface) DeleteUser(c *gin.Context) {
	err := uci.service.DeleteUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso"})
}
