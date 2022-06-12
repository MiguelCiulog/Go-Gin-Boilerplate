package routes

import (
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouteUser(parent *gin.RouterGroup) *gin.RouterGroup {

	userGroup := parent.Group("users")
	{
		user := new(controllers.UserController)
		userGroup.GET("/", user.RetrieveAllUsers)
		userGroup.GET("/:id", user.GetUserById)
	}
	return userGroup
}
