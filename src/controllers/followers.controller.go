package controllers

import (
	internal "api/internal/db"
	"api/src/auth"
	followersRepository "api/src/repositories/followers"
	"api/src/response"
	"api/src/validation"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func FollowerUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	userId := params["userId"]

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusBadRequest, err)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Followers(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	userId := params["userId"]

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
		return
	}

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Unfollow(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func GetFollow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId := params["userId"]

	db, err := internal.PrismaClientDB()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	repo := followersRepository.FollowersRepository(db)
	followers, err := repo.GetFollow(userId)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	followId := params["followId"]

	db, err := internal.PrismaClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := followersRepository.FollowersRepository(db)

	followers, err := repo.GetFollowers(followId)

	validation.Err(w, http.StatusInternalServerError, err)

	response.JSON(w, http.StatusOK, followers)
}
