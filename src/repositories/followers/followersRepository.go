package followersRepository

import (
	"database/sql"
	"log"
)

type User struct {
	db *sql.DB
}

func FollowersRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) Followers(userId, followersId uint64) error {
	statement, err := u.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}

	defer func(statements *sql.Stmt) {
		err := statements.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)
	if _, err = statement.Exec(userId, followersId); err != nil {
		return err
	}
	return nil
}
