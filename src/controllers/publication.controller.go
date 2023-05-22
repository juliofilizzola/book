package controllers

import (
	db2 "api/db"
	"api/src/auth"
	"api/src/models"
	publication2 "api/src/repositories/publication"
	"api/src/response"
	"api/src/validation"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
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

func GetMyPublications(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}
	db, err := db2.Connection()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublicationByUser(userId)
	if err != nil {
		response.Err(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusOK, data)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	validation.Err(w, http.StatusBadRequest, err)

	body, err := io.ReadAll(r.Body)
	validation.Err(w, http.StatusUnprocessableEntity, err)

	var publication models.Publication

	if err = json.Unmarshal(body, &publication); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db2.Connection()
	validation.Err(w, http.StatusInternalServerError, err)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	repo := publication2.PublicationsRepository(db)

	err = repo.UpdatePublication(id, userId, publication)
	validation.Err(w, http.StatusInternalServerError, err)

	response.JSON(w, http.StatusNoContent, nil)
}