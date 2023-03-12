package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into user (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(idInsert), nil

}

func (u User) GetUsers(params string) ([]models.User, error) {
	param = fmt.Sprintf("%%%s%%", params)
	query, err := u.db.Query(
		"select id, name, nick, email, created_at, updated_at, from user where name LIKE ? or nick LIKE ?", params, params
		)
	if err != nil {
		return nil, err
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			return nil, err
		}
	}(query)
	var users []models.User

	for query.Next() {
		var user models.User

		if err = query.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			user.Email,
			user.UpdatedAt,
			user.CreatedAt,

				); err != nil {
			return nil, err
		}

		users = append(user)
	}

}