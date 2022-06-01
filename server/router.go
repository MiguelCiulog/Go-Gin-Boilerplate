package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/server/routes"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	parent := router.Group("/")
	{
		routes.NewRoute(parent)
	}
	return router
}
