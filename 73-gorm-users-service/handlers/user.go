package handlers

import (
	"log/slog"
	"strconv"
	"time"
	"user-service/database"
	"user-service/models"
	"user-service/security"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetAllByLimit(ctx *gin.Context)
}

type UserHandler struct {
	UsersDB   *database.UserDB
	JWTSecret string
}

func NewUserHandler(userDb *database.UserDB, jwtSecret string) *UserHandler {
	return &UserHandler{UsersDB: userDb, JWTSecret: jwtSecret}
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.Bind(user)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	err = user.Validate()
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	user.Status = "active"
	user.LastUpdated = uint64(time.Now().Unix())
	hash, err := security.HashPassword(user.Password)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	user.Password = hash

	user, err = uh.UsersDB.Create(user)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	user.Password = "*************"

	ctx.JSON(201, user)

}

func (uh *UserHandler) Login(ctx *gin.Context) {
	userLogin := new(models.UserLogin)
	err := ctx.Bind(userLogin)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}
	user, err := uh.UsersDB.GetUserByEmailWithPassword(userLogin.Email)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	err = security.VerifyPassword(userLogin.Password, user.Password)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Invalid email or passowrd")
		ctx.Abort()
		return
	}

	//ctx.String(201, "User successfully login")

	// Here need to return the JWT token

	token, err := security.GenerateJWT(user.Name, user.Email, uh.JWTSecret)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Error in issuing the token")
		ctx.Abort()
		return
	}

	ctx.JSON(201, map[string]string{"token": token})
	//ctx.JSON(201,gin.H{"token":token})
}

func (uh *UserHandler) GetAll(ctx *gin.Context) {
	users, err := uh.UsersDB.GetAll()
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}
	ctx.JSON(200, users)
}

func (uh *UserHandler) GetAllByLimit(ctx *gin.Context) {

	limit, ok := ctx.Params.Get("limit")

	if !ok {
		slog.Error("invalid limit value")
		ctx.String(400, "limit param is wrong")
		ctx.Abort()
		return
	}

	_limit, err := strconv.Atoi(limit)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	offset, ok := ctx.Params.Get("offset")

	if !ok {
		slog.Error("invalid offset value")
		ctx.String(400, "offset param is wrong")
		ctx.Abort()
		return
	}

	_offset, err := strconv.Atoi(offset)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	users, err := uh.UsersDB.GetAllByLimit(_limit, _offset)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}
	ctx.JSON(200, users)
}
