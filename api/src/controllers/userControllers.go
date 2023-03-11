package UserControllers

import (
	"api/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
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

	db2, err := db.Connection()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	repo := repositories.UserRepository(db2)
	user.ID, err = repo.Create(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		Nick      string    `json:"nick"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        user.ID,
		Nick:      user.Nick,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creater User"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creater User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creater User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creater User"))
}
