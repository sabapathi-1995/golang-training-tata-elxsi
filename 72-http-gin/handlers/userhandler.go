package handlers

import (
	"bufio"
	"demo/fileops"
	"demo/models"
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type IUser interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetFileName() string
	//GetUsers(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	FileName string
}

func NewUserHandler(fileName string) IUser {
	return &UserHandler{FileName: fileName}
}

func (uh *UserHandler) GetFileName() string {
	return uh.FileName
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	// from request body I get the data in the form of json

	user := new(models.User)

	if err := c.Bind(user); err != nil { // automatically binds the data from http request into the object.. Deserialization happnes easily
		// Deserialize
		c.String(http.StatusBadRequest, err.Error())
		c.Abort() // dont move next , just stop here
		return
	}

	user.Id = int(rand.Int64())
	user.Status = "active"
	user.LastModified = time.Now().Unix() // unix time

	// return back to the response

	if bytes, err := json.Marshal(user); err != nil {
		slog.Error(err.Error())
		c.String(400, "Opps.Something went wrong!")
		c.Abort()
		return
	} else {
		//	err := fileops.SaveToFile(uh.FileName, bytes)
		fileops.ChanData <- bytes // sending data to channel
		if err != nil {
			slog.Error(err.Error())
			c.String(400, "Opps.Something went wrong!")
			c.Abort()
			return
		}
		c.JSON(201, user)
		// c.Request.Response.Write(bytes)
		// c.Request.Response.WriteHeader(201)

	}

}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	users := make([]models.User, 0)

	file, err := os.Open(uh.FileName)
	if err != nil {
		slog.Error(err.Error())
		c.String(400, "Opps.Something went wrong!")
		c.Abort()
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		user := models.User{}
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			slog.Error(err.Error())
			c.String(400, "Opps.Something went wrong!")
			c.Abort()
			return
		}
		users = append(users, user)
	}

	c.JSON(200, users)

}

// Take the data and store it in a file as users.db
// make sure each user is wriiten in a new line
// write it as a json string
