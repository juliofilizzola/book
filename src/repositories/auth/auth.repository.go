package authRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Auth struct {
	db *sql.DB
}

func AuthRepository(db *sql.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (a Auth) SearchPassword(userId uint64) (string, error) {
	var (
		password = ""
	)

	param := fmt.Sprintf("%%%s%%", userId)
	query, err := a.db.Query(
		"select password from user where user.id = ? ", param,
	)

	if err != nil {
		return password, err
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	if query.Next() {
		if err = query.Scan(
			&password,
		); err != nil {
			return password, err
		}
	} else {
		return password, errors.New("not found")
	}
	return password, nil
}

func (a Auth) UpdatePassword(newPassword []byte, userId uint64) error {
	statement, err := a.db.Prepare("update user set password = ? where user.id = ?")

	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	if _, err := statement.Exec(newPassword, userId); err != nil {
		return err
	}

	return nil
}
