package routes

import (
	"github.com/eron97/inject.git/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controllers.ControllerInterface) {
	r.GET("/getAllUsers", controller.ReadlAllUsers)
	r.GET("/getUser/:id", controller.ReadUser)
	r.POST("/createUser", controller.CreateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)
}
