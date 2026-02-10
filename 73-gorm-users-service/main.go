package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime"
	"user-service/database"
	"user-service/handlers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

var (
	PORT, DB_URL string
)

func init() {

	PORT = os.Getenv("APP_PORT")
	DB_URL = os.Getenv("DB_URL")
	// if PORT == "" {
	// 	PORT = "8083"
	// }
	// "host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
}

func main() {

	//args := os.Args

	if PORT == "" {
		flag.StringVar(&PORT, "port", "8083", "provide the port. If port is not envirement and not given thru options , it takes the default port")
	}

	if DB_URL == "" {
		flag.StringVar(&DB_URL, "db", `host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "give postgres db connection")
	}

	flag.Parse()

	router := gin.Default()

	router.GET("/", handlers.Root)
	router.GET("/ping", handlers.Ping)
	router.GET("/health", handlers.Health)

	db, err := database.ConnectDb(DB_URL)
	if err != nil {
		panic(err.Error())
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err.Error())
	}

	jwtSecret := "my_jwt_secret"

	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB, jwtSecret)

	userRouter := router.Group("/v1/public")

	userRouter.POST("/user", userHandler.Create)
	userRouter.POST("/user/login", userHandler.Login)

	privateRouter := router.Group("/v1/private", middleware.JWTAuth(jwtSecret)) // If I want to use any kind of private uri, it has to be authorized with the token

	productHandler := handlers.NewProductHandler(jwtSecret)
	privateRouter.POST("/product", func(ctx *gin.Context) {
	}, productHandler.Create)

	privateRouter.GET("/user/all", userHandler.GetAll)
	privateRouter.GET("/user/all/:limit/:offset", userHandler.GetAllByLimit)
	slog.Info("The application is listening on port:" + PORT)
	if err := router.Run("0.0.0.0:" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

}
