package controllers

import (
	"api/cmd/auth"
	"api/cmd/repositories/followers"
	"api/cmd/response"
	"api/cmd/validation"
	internal "api/internal/db"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func FollowerUser(w http.ResponseWriter, r *http.Request) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	followerId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	userId := params["userId"]

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
		return
	}

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusBadRequest, err)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Followers(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	followerId, err := auth.GetUserId(r)

	validation.Err(w, http.StatusUnauthorized, err)

	params := mux.Vars(r)

	userId := params["userId"]

	if userId == followerId {
		response.Err(w, http.StatusForbidden, errors.New("cannot follower you"))
		return
	}

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := followersRepository.FollowersRepository(db)

	if err = repo.Unfollow(userId, followerId); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func GetFollow(w http.ResponseWriter, r *http.Request) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	params := mux.Vars(r)

	userId := params["userId"]

	db, err := internal.ClientDB()

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
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	params := mux.Vars(r)

	followId := params["followId"]

	db, err := internal.ClientDB()

	validation.Err(w, http.StatusInternalServerError, err)

	repo := followersRepository.FollowersRepository(db)

	followers, err := repo.GetFollowers(followId)

	validation.Err(w, http.StatusInternalServerError, err)

	response.JSON(w, http.StatusOK, followers)
}
