package followersRepository

import (
	"api/src/models"
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

func (u User) Unfollow(userId, followerId uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}

	defer func(statements *sql.Stmt) {
		err := statements.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (u User) GetFollow(userId uint64) ([]models.User, error) {
	row, err := u.db.Query(`
		select user.id, user.name, user.email, user.nick, user.created_at
		from USER inner join FOLLOWERS F on user.id = F.follower_id where user.id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var users []models.User
	for row.Next() {
		var user models.User
		if err = row.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u User) GetFollowers(followId uint64) ([]models.User, error) {
	row, err := u.db.Query(`
		select user.id, user.name, user.email, user.nick, user.created_at
		from USER inner join FOLLOWERS F on user.id = F.follower_id where F.follower_id = ?
	`, followId)
	if err != nil {
		return nil, err
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var users []models.User
	for row.Next() {
		var user models.User
		if err = row.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
