package routes

import (
	"github.com/eron97/inject.git/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller controllers.ControllerInterface) {

	r.POST("/createUser", controller.CreateUser)

}
