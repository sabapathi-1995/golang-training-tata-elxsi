package app

import (
	"net/http"

	"gin-gorm-postgres-crud-itest/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

	h := handlers.NewUserHandler(db)

	v1 := r.Group("/v1")
	{
		v1.POST("/users", h.Create)
		v1.GET("/users/:id", h.Get)
		v1.GET("/users", h.List)
		v1.PUT("/users/:id", h.Update)
		v1.DELETE("/users/:id", h.Delete)
	}

	return r
}
