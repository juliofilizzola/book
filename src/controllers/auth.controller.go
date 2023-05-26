package controllers

import (
	internal "api/internal/db"
	"api/src/auth"
	"api/src/models"
	authRepository "api/src/repositories/auth"
	"api/src/response"
	"api/src/validation"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	userId := params["userId"]

	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if userIDToken != userId {
		response.Err(w, http.StatusUnauthorized, errors.New("cannot update user"))
		return
	}

	body, err := io.ReadAll(r.Body)

	var password models.Password

	if err = json.Unmarshal(body, &password); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := authRepository.AuthRepository(db)

	passwordDb, err := repo.SearchPassword(userId)

	validation.Err(w, http.StatusInternalServerError, err)

	if err = auth.ValidPassword(passwordDb, password.CurrentPassword); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	newPassword, err := auth.Hash(password.NewPassword)

	validation.Err(w, http.StatusBadRequest, err)

	if err = repo.UpdatePassword(newPassword, userId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}
