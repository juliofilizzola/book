package loginRepository

import (
	"api/prisma/db"
	"api/src/models"
	"context"
)

type User struct {
	db *db.PrismaClient
}

func LoginRepository(db *db.PrismaClient) *User {
	return &User{db}
}

func (u User) SEARCH(params string) (models.User, error) {
	ctx := context.Background()

	query, err := u.db.User.FindFirst(
		db.User.Email.Set(params),
	).Exec(ctx)

	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:       query.ID,
		Name:     query.Name,
		Email:    query.Email,
		Nick:     query.Nick,
		Password: query.Password,
	}, err
	//param := fmt.Sprintf("%%%s%%", params)
	//query, err := u.db.Query(
	//	"select id, password, email, created_at from user where email LIKE ? ", param,
	//)
	//
	//if err != nil {
	//	return models.User{}, err
	//}
	//
	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(query)
	//
	//var user models.User
	//if query.Next() {
	//	if err = query.Scan(
	//		&user.ID,
	//		&user.Password,
	//		&user.Email,
	//		&user.Nick,
	//	); err != nil {
	//		return models.User{}, err
	//	}
	//} else {
	//	return models.User{}, errors.New("not found")
	//}
	//return user, nil
}
