package controllers

import (
	"net/http"

	service "github.com/eron97/inject.git/services"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	CreateUser(c *gin.Context)
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

func (uci *useControllerInterface) CreateUser(c *gin.Context) {
	uci.service.CreateUserServices()

	c.JSON(http.StatusOK, gin.H{"message": "Usuário criado com sucesso!"})
}
