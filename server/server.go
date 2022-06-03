package server

// import "github.com/MiguelCiulog/Go-Gin-Boilerplate/config"

func Init() {
	// config := config.GetConfig()
	r := NewRouter()
	// r.Run(config.GetString("server.port"))
	r.Run(":8080")
}
