package main

import (
	// "flag"
	"fmt"
	// "os"

	"github.com/MiguelCiulog/Go-Gin-Boilerplate/config"
	// "github.com/MiguelCiulog/Go-Gin-Boilerplate/db"
	// "github.com/MiguelCiulog/Go-Gin-Boilerplate/server"
)

func main() {
	// environment := flag.String("e", "development", "")
	// flag.Usage = func() {
	// 	fmt.Println("Usage: server -e {mode}")
	// 	os.Exit(1)
	// }
	// flag.Parse()
	// config.Init(*environment)
	config.Init("development")
	fmt.Println(config.GetConfig().GetString("db.database"))
	// db.Init()
	// server.Init()
}
