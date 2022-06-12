package main

import (
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/config"
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/db"
	"github.com/MiguelCiulog/Go-Gin-Boilerplate/server"
)

func main() {
	config.Init("development")
	db.Init()
	server.Init()
}
