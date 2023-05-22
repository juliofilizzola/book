package controllers

import (
	"api/db"
	"api/src/auth"
	"api/src/models"
	loginRepository "api/src/repositories/login"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
	defer func(db2 *sql.DB) {
		err := db2.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db2)

	repo := loginRepository.LoginRepository(db2)

	userSearch, err := repo.SEARCH(user.Email)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = auth.ValidPassword(userSearch.Password, user.Password); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}
	var responseToken models.Login
	token, err := auth.GenerateToken(userSearch.ID)

	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}
	responseToken.Token = token
	response.JSON(w, http.StatusAccepted, responseToken)
}
