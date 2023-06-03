package followersRepository

import (
	"api/cmd/models"
	"database/sql"
	"log"
)

type User struct {
	db *sql.DB
}

func FollowersRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) Followers(userId, followersId string) error {
	//ctx := context.Background()
	//
	//_, err := u.db.Followers.CreateOne(
	//	db.Followers.Follower.Link(
	//		db.User.ID.Equals(userId),
	//	),
	//	db.Followers.Following.Link(
	//		db.User.ID.Equals(followersId),
	//	),
	//).Exec(ctx)

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

func (u User) Unfollow(userId, followerId string) error {
	//ctx := context.Background()
	//
	//_, err := u.db.Followers.FindUnique(
	//	db.Followers.FollowerIDFollowingID(
	//		db.Followers.FollowerID.Equals(userId),
	//		db.Followers.FollowingID.Equals(userId),
	//	),
	//).Delete().Exec(ctx)
	//
	//if err != nil {
	//	return err
	//}
	//
	//return nil

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

func (u User) GetFollow(userId string) ([]models.User, error) {
	//ctx := context.Background()
	//var (
	//	users []models.User
	//)
	//err := u.db.Prisma.QueryRaw("select u.id, u.name, u.email, u.nick, u.created_at from USER u inner join FOLLOWERS F on u.id = F.follower_id where u.id = ?", userId).Exec(ctx, &users)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return users, nil

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

func (u User) GetFollowers(followId string) ([]models.User, error) {

	//ctx := context.Background()
	//var (
	//	users []models.User
	//)
	//err := u.db.Prisma.QueryRaw("select u.id, u.name, u.email, u.nick, u.created_at from USER u inner join FOLLOWERS F on u.id = F.follower_id where F.follower_id = ?", followId).Exec(ctx, &users)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return users, nil
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
