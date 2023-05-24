package controllers

import (
	db2 "api/db"
	"api/src/auth"
	"api/src/models"
	authRepository "api/src/repositories/auth"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserId(r)

	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)

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

	db, err := db2.Connection()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	repo := authRepository.AuthRepository(db)

	passwordDb, err := repo.SearchPassword(userId)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = auth.ValidPassword(passwordDb, password.CurrentPassword); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	newPassword, err := auth.Hash(password.NewPassword)

	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePassword(newPassword, userId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}
