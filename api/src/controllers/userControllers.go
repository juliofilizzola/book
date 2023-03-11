package UserControllers

import (
	"api/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}
	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		log.Fatal(error)
	}

	db2, err := db.Connection()
	if err != nil {
		log.Fatal(error)
	}
	repo := repositories.UserRepository(db2)
	create, err := repo.Create(user)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id insert: %d", create)))
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
