package users

import (
	"api/prisma/db"
	"api/src/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type User struct {
	db *db.PrismaClient
}

func UserRepository(db *db.PrismaClient) *User {
	return &User{
		db: db,
	}
}

func (u User) Create(user models.User) (string, error) {
	ctx := context.Background()

	result, err := u.db.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Nick.Set(user.Nick),
		db.User.Password.Set(user.Password),
		db.User.Email.Set(user.Email),
	).Exec(ctx)

	if err != nil {
		return "", err
	}

	return result.ID, nil

	//statement, err := u.db.Prepare("insert into user (name, nick, email, password) values (?, ?, ?, ?)")
	//if err != nil {
	//	return 0, err
	//}
	//defer func(statement *sql.Stmt) {
	//	err := statement.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(statement)
	//
	//insert, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	//
	//if err != nil {
	//	return 0, err
	//}
	//
	//idInsert, err := insert.LastInsertId()
	//if err != nil {
	//	return 0, err
	//}
	//
	//return uint64(idInsert), nil

}

func (u User) GetUsers(params string) ([]models.User, error) {
	ctx := context.Background()

	param := fmt.Sprintf("%%%s%%", params)

	var (
		users []models.User
	)

	err := u.db.Prisma.QueryRaw("select id, name, nick, email, created_at from user where name LIKE ? or nick LIKE ?", param, param).Exec(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
	//query, err := u.db.Query(
	//	"select id, name, nick, email, created_at from user where name LIKE ? or nick LIKE ?", param, param,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(query)
	//var users []models.User
	//
	//for query.Next() {
	//	var user models.User
	//	if err = query.Scan(
	//		&user.ID,
	//		&user.Name,
	//		&user.Nick,
	//		&user.Email,
	//		&user.CreatedAt,
	//	); err != nil {
	//		return nil, err
	//	}
	//
	//	users = append(users, user)
	//}
	//return users, nil
}

func (u User) GetUser(ID uint64) (models.User, error) {

	getUser, err := u.db.Query("select id, name, nick, email, created_at from user where id = ? ", ID)

	if err != nil {
		fmt.Println("!")
		return models.User{}, err
	}

	defer func(getUser *sql.Rows) {
		err := getUser.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(getUser)

	var user models.User

	if getUser.Next() {
		if err = getUser.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, errors.New("not found")
	}

	return user, nil
}

func (u User) UpdatedUser(id uint64, body models.User) error {
	statement, err := u.db.Prepare("update user set name = ?, nick = ?, email = ? where id = ? ")

	if err != nil {
		return err
	}
	defer func(statements *sql.Stmt) {
		err := statements.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	if _, err = statement.Exec(body.Name, body.Nick, body.Email, id); err != nil {
		return err
	}
	return nil
}

func (u User) DeleteUser(id uint64) error {
	statement, err := u.db.Prepare("delete from user where id = ?")
	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)
	if _, err = statement.Exec(id); err != nil {
		return err
	}
	return nil
}
