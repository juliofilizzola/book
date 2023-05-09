package loginRepository

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type User struct {
	db *sql.DB
}

func LoginRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) SEARCH(params string) (models.User, error) {
	param := fmt.Sprintf("%%%s%%", params)
	query, err := u.db.Query(
		"select id, password, email, created_at from user where email LIKE ? ", param,
	)

	if err != nil {
		return models.User{}, err
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	var user models.User
	if query.Next() {
		if err = query.Scan(
			&user.ID,
			&user.Password,
			&user.Email,
			&user.Nick,
		); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, errors.New("not found")
	}
	return user, nil
}
