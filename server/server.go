package server

// import "github.com/MiguelCiulog/Go-Gin-Boilerplate/config"

func Init() {
	// TODO: configure router to use enviroment
	r := NewRouter()
	r.Run(":8080")
}
