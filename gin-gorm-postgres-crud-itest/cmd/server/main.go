package main

import (
	"log"
	"os"

	"gin-gorm-postgres-crud-itest/internal/app"
	"gin-gorm-postgres-crud-itest/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// Example DSN for local Postgres.
		dsn = "host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	}

	gdb, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	if err := db.Migrate(gdb); err != nil {
		log.Fatalf("migrate failed: %v", err)
	}

	r := app.NewRouter(gdb)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
