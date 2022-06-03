package routes

import (
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func NewRoute(parent *gin.RouterGroup) *gin.RouterGroup {

	userGroup := parent.Group("users")
	{
		user := new(controllers.UserController)
		userGroup.GET("/:id", user.Retrieve)
	}
	return userGroup
}
