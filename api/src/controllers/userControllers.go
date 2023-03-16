package UserControllers

import (
	"api/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareData(false); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db2, err := db.Connection()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db2 *sql.DB) {
		err := db2.Close()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db2)
	repo := repositories.UserRepository(db2)
	user.ID, err = repo.Create(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, struct {
		ID        uint64    `json:"id"`
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
	db2, err := db.Connection()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db2 *sql.DB) {
		err := db2.Close()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db2)

	repo := repositories.UserRepository(db2)
	users, err := repo.GetUsers(params)
	if err != nil {
		response.Err(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
	}
	db2, err := db.Connection()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db2 *sql.DB) {
		err := db2.Close()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db2)

	repo := repositories.UserRepository(db2)
	users, err := repo.GetUser(ID)
	if err != nil {
		response.Err(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareData(true); err != nil {
		response.Err(w, http.StatusConflict, err)
		return
	}

	db2, err := db.Connection()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db2 *sql.DB) {
		err := db2.Close()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db2)
	repo := repositories.UserRepository(db2)

	if err := repo.UpdatedUser(ID, user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creater User"))
}
