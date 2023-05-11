package FollowersController

import (
	db2 "api/db"
	"api/src/auth"
	followersRepository "api/src/repositories/followers"
	"errors"

	//"api/src/auth"
	//followersRepository "api/src/repositories/followers"
	"api/src/response"
	"database/sql"
	//"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FollowerUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId(r)

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

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
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
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Followers(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId(r)

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

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
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
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Unfollow(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)

}

func GetFollow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
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
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}(db)

	repo := followersRepository.FollowersRepository(db)
	followers, err := repo.GetFollow(userId)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}
