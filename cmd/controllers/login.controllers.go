package controllers

import (
	auth2 "api/cmd/auth"
	models2 "api/cmd/models"
	"api/cmd/repositories/login"
	"api/cmd/response"
	internal "api/internal/db"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models2.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.ClientDB()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	repo := loginRepository.LoginRepository(db)

	userSearch, err := repo.SEARCH(user.Email)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = auth2.ValidPassword(userSearch.Password, user.Password); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}
	var responseToken models2.Login

	token, err := auth2.GenerateToken(userSearch.ID)

	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	responseToken.Token = token

	response.JSON(w, http.StatusAccepted, responseToken)
}
