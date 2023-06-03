package controllers

import (
	"api/cmd/auth"
	"api/cmd/models"
	"api/cmd/repositories/users"
	"api/cmd/response"
	"api/cmd/validation"
	internal "api/internal/db"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	validation.Err(w, http.StatusUnprocessableEntity, err)

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareData(false); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := users.UserRepository(db)

	user.ID, err = repo.Create(user)

	validation.Err(w, http.StatusInternalServerError, err)

	response.JSON(w, http.StatusCreated, struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Nick      string    `json:"nick"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        user.ID,
		Email:     user.Email,
		Nick:      user.Nick,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.URL.Query().Get("user"))

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := users.UserRepository(db)

	usersResponse, err := repo.GetUsers(params)

	validation.Err(w, http.StatusNotFound, err)

	response.JSON(w, http.StatusOK, usersResponse)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["id"]

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := users.UserRepository(db)

	usersResponse, err := repo.GetUser(userId)

	validation.Err(w, http.StatusNotFound, err)

	response.JSON(w, http.StatusOK, usersResponse)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]

	if valid := auth.ValidUser(r, userId); !valid {
		response.Err(w, http.StatusUnauthorized, errors.New("cannot updated other user"))
		return
	}

	body, err := io.ReadAll(r.Body)

	validation.Err(w, http.StatusInternalServerError, err)

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareData(true); err != nil {
		response.Err(w, http.StatusConflict, err)
		return
	}

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := users.UserRepository(db)

	if err := repo.UpdatedUser(userId, user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["id"]

	if valid := auth.ValidUser(r, userId); !valid {
		response.Err(w, http.StatusUnauthorized, errors.New("cannot deleted other user"))
		return
	}

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := users.UserRepository(db)

	if err = repo.DeleteUser(userId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
