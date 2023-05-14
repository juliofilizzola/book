package controllers

import (
	db2 "api/db"
	"api/src/auth"
	"api/src/models"
	publication2 "api/src/repositories/publication"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)

	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication

	if err = json.Unmarshal(body, &publication); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthId = strconv.FormatUint(userId, 10)
	publication.Likes = strconv.Itoa(0)
	db, err := db2.Connection()

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
	}(db)

	repo := publication2.PublicationsRepository(db)

	publication.ID, err = repo.Create(publication)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, struct {
		models.Publication
	}{
		publication,
	})
}
