package controllers

import (
	auth2 "api/cmd/auth"
	"api/cmd/models"
	"api/cmd/repositories/auth"
	"api/cmd/response"
	"api/cmd/validation"
	internal "api/internal/db"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"sync"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	userIDToken, err := auth2.GetUserId(r)

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

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := authRepository.AuthRepository(db)

	passwordDb, err := repo.SearchPassword(userId)

	validation.Err(w, http.StatusInternalServerError, err)

	if err = auth2.ValidPassword(passwordDb, password.CurrentPassword); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	newPassword, err := auth2.Hash(password.NewPassword)

	validation.Err(w, http.StatusBadRequest, err)

	if err = repo.UpdatePassword(newPassword, userId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}
