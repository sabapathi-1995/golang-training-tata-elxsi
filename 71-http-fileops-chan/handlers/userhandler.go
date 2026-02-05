package handlers

import (
	"demo/fileops"
	"demo/models"
	"encoding/json"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"
)

type IUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
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

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// from request body I get the data in the form of json

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte("Opps.Something went wrong!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := new(models.User)

	if err := json.Unmarshal(data, user); err != nil {
		slog.Error(err.Error())
		w.Write([]byte("Opps.Something went wrong!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Deserialize

	user.Id = int(rand.Int64())
	user.Status = "active"
	user.LastModified = time.Now().Unix() // unix time

	// return back to the response

	if bytes, err := json.Marshal(user); err != nil {
		slog.Error(err.Error())
		w.Write([]byte("Opps.Something went wrong!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {

		//	err := fileops.SaveToFile(uh.FileName, bytes)
		fileops.ChanData <- bytes // sending data to channel

		if err != nil {
			slog.Error(err.Error())
			w.Write([]byte("Opps.Something went wrong!"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(bytes)
		w.WriteHeader(201)
		//w.WriteHeader(http.StatusCreated)
	}

}

// Take the data and store it in a file as users.db
// make sure each user is wriiten in a new line
// write it as a json string
