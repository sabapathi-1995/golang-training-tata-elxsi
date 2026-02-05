package main

import (
	"demo/fileops"
	"demo/handlers"
	"flag"
	"log/slog"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func init() {

	PORT = os.Getenv("APP_PORT")
	// if PORT == "" {
	// 	PORT = "8083"
	// }
}

func main() {

	//args := os.Args

	if PORT == "" {
		flag.StringVar(&PORT, "port", "8083", "provide the port. If port is not envirement and not given thru options , it takes the default port")
		flag.Parse()
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome Gin package")
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	router.GET("/health", Health)

	userhandler := handlers.NewUserHandler("users.txt")

	fileops.Init(userhandler.GetFileName())

	router.POST("/user", userhandler.CreateUser)
	router.GET("/users", userhandler.GetUsers)

	slog.Info("The application is listening on port:" + PORT)
	if err := router.Run("0.0.0.0:" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	// all network interfaces:8083
}

func Health(ctx *gin.Context) {
	ctx.String(200, "ok")
}

// Cloud Native Applications
// Go does not require a webserver like Apache Tomcat/IIS
// It has self hosted applications
// The binary itself contains a packed webserver, the moment it see http package

// listener
// Router
// ResponseWriter
// Request <- All the request data , body, headers, params, authentication heardss

// docker pull postgres
