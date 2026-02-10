package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type IProduct interface {
	Create(ctx *gin.Context)
}

type ProductHandler struct {
	// UsersDB   *database.UserDB
	JWTSecret string
}

func NewProductHandler(jwtSecret string) *ProductHandler {
	return &ProductHandler{JWTSecret: jwtSecret}
}

func (uh *ProductHandler) Create(ctx *gin.Context) {
	product := make(map[string]any)

	err := ctx.Bind(&product)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	ctx.JSON(201, product)

}
