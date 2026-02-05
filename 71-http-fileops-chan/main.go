package main

import (
	"demo/fileops"
	"demo/handlers"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"runtime"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/health", Health)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	userhandler := handlers.NewUserHandler("users.txt")

	fileops.Init(userhandler.GetFileName())

	http.HandleFunc("/user", userhandler.CreateUser) // create user

	slog.Info("The application is listening on port:" + PORT)
	if err := http.ListenAndServe("0.0.0.0:"+PORT, nil); err != nil {
		slog.Error(err.Error())
		runtime.Goexit()
	}

	// all network interfaces:8083
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	w.WriteHeader(http.StatusOK)
}

// Cloud Native Applications
// Go does not require a webserver like Apache Tomcat/IIS
// It has self hosted applications
// The binary itself contains a packed webserver, the moment it see http package

// listener
// Router
// ResponseWriter
// Request <- All the request data , body, headers, params, authentication heardss
