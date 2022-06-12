package controllers

import (
	"net/http"
	"strconv"

	"github.com/MiguelCiulog/Go-Gin-Boilerplate/models"
	"github.com/gin-gonic/gin"
)

// Used to access Retrieve method
type UserController struct{}

var userModel = new(models.User)

func (u UserController) RetrieveAllUsers(c *gin.Context) {
	users, err := userModel.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

func (u UserController) GetUserById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error to retrieve user",
				"error": "Error parsing id parameter"})
		c.Abort()
	}
	users, err := userModel.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
