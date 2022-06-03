package server

import (
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/server/routes"
	"github.com/gin-gonic/gin"
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
