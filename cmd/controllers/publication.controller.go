package controllers

import (
	"api/cmd/auth"
	"api/cmd/models"
	publication2 "api/cmd/repositories/publication"
	"api/cmd/response"
	"api/cmd/validation"
	internal "api/internal/db"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	body, err := io.ReadAll(r.Body)

	validation.Err(w, http.StatusUnprocessableEntity, err)

	var publication models.Publication

	if err = json.Unmarshal(body, &publication); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthId = userId
	publication.Likes = 0

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusBadRequest, err)

	repo := publication2.PublicationsRepository(db)

	publication.ID, err = repo.Create(publication)

	validation.Err(w, http.StatusUnprocessableEntity, err)

	response.JSON(w, http.StatusCreated, struct {
		models.Publication
	}{
		publication,
	})
}

func GetMyPublications(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusBadRequest, err)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublicationByUser(userId)

	validation.Err(w, http.StatusNotFound, err)

	response.JSON(w, http.StatusOK, data)
}

func GetPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusBadRequest, err)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublication(id)

	validation.Err(w, http.StatusNotFound, err)

	response.JSON(w, http.StatusOK, data)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	_, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	id := params["id"]

	validation.Err(w, http.StatusBadRequest, err)

	body, err := io.ReadAll(r.Body)

	validation.Err(w, http.StatusUnprocessableEntity, err)

	var publication models.Publication

	if err = json.Unmarshal(body, &publication); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := publication2.PublicationsRepository(db)

	err = repo.UpdatePublication(id, publication)

	validation.Err(w, http.StatusInternalServerError, err)

	response.JSON(w, http.StatusNoContent, nil)
}

func GetAllPublication(w http.ResponseWriter, r *http.Request) {
	_, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublications()

	validation.Err(w, http.StatusNotFound, err)

	response.JSON(w, http.StatusOK, data)
}

func DeletedPublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	id := params["id"]

	validation.Err(w, http.StatusBadRequest, err)

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := publication2.PublicationsRepository(db)

	dbPublication, err := repo.GetPublication(id)

	validation.Err(w, http.StatusBadRequest, err)

	authId := dbPublication.AuthId

	if authId != userId {
		response.Err(w, http.StatusBadRequest, errors.New("token invalid"))
	}

	err = repo.DeletedPublication(id)

	validation.Err(w, http.StatusBadRequest, err)

	response.JSON(w, http.StatusNoContent, nil)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	body, err := io.ReadAll(r.Body)

	validation.Err(w, http.StatusUnprocessableEntity, err)

	var like models.Like

	if err = json.Unmarshal(body, &like); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublication(id)

	validation.Err(w, http.StatusNotFound, err)

	likes := data.Likes

	validation.Err(w, http.StatusNotFound, err)

	var (
		totalLikes = likes + like.Like
	)

	err = repo.LikePublication(id, totalLikes)

	validation.Err(w, http.StatusBadRequest, err)

	response.JSON(w, http.StatusNoContent, nil)
}

func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	body, err := io.ReadAll(r.Body)

	validation.Err(w, http.StatusUnprocessableEntity, err)

	var dislike models.Like

	if err = json.Unmarshal(body, &dislike); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := publication2.PublicationsRepository(db)

	data, err := repo.GetPublication(id)

	validation.Err(w, http.StatusNotFound, err)

	likes := data.Likes

	validation.Err(w, http.StatusNotFound, err)

	var (
		totalLikes = likes - dislike.Like
	)

	err = repo.LikePublication(id, int(totalLikes))

	validation.Err(w, http.StatusBadRequest, err)

	response.JSON(w, http.StatusNoContent, nil)
}
