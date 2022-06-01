package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/controllers"
)

func NewRoute(parent *gin.RouterGroup) *gin.RouterGroup {

	userGroup := parent.Group("users")
	{
		user := new(controllers.UserController)
		userGroup.GET("/:id", user.Retrieve)
	}
	return userGroup
}
