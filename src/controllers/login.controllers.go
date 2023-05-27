package controllers

import (
	internal "api/internal/db"
	"api/src/auth"
	"api/src/models"
	loginRepository "api/src/repositories/login"
	"api/src/response"
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

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.PrismaClientDB()

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
