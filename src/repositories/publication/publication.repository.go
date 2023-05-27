package publication

import (
	"api/prisma/db"
	"api/src/models"
	"context"
)

type Publication struct {
	db *db.PrismaClient
}

func PublicationsRepository(db *db.PrismaClient) *Publication {
	return &Publication{
		db: db,
	}
}

func (p Publication) Create(Publication models.Publication) (string, error) {
	ctx := context.Background()

	create, err := p.db.Publication.CreateOne(
		db.Publication.Title.Set(Publication.Title),
		db.Publication.Description.Set(Publication.Description),
		db.Publication.Content.Set(Publication.Content),
		db.Publication.Like.Set(Publication.Likes),
		db.Publication.Auth.Link(
			db.User.ID.Equals(Publication.AuthId),
		),
	).Exec(ctx)

	if err != nil {
		return "", err
	}

	return create.ID, nil
	//statement, err := p.db.Prepare("insert into PUBLICATION (title, auth_id, description, content, `like`) values (?, ?, ?, ?, ?)")
	//if err != nil {
	//	return 0, err
	//}
	//
	//defer func(statement *sql.Stmt) {
	//	err := statement.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(statement)
	//
	//insert, err := statement.Exec(Publication.Title, Publication.AuthId, Publication.Description, Publication.Content, Publication.Likes)
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

func (p Publication) GetPublicationByUser(id string) ([]models.PublicationReturn, error) {
	ctx := context.Background()

	var (
		publications []models.PublicationReturn
	)

	err := p.db.Prisma.QueryRaw(`select
		P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER U on U.id = P.auth_id where u.id = ?`, id).Exec(ctx, &publications)

	if err != nil {
		return nil, err
	}

	return publications, nil

	//query, err := p.db.Query(`select
	//	P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
	//	from PUBLICATION as P inner join USER U on U.id = P.auth_id where u.id = ?`, id)
	//
	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(query)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//var publications []models.PublicationReturn
	//for query.Next() {
	//	var publication models.PublicationReturn
	//	if err = query.Scan(
	//		&publication.ID,
	//		&publication.Title,
	//		&publication.Description,
	//		&publication.Content,
	//		&publication.Likes,
	//		&publication.CreatedAt,
	//		&publication.AuthNick,
	//		&publication.AuthEmail,
	//	); err != nil {
	//		return nil, err
	//	}
	//	publications = append(publications, publication)
	//}
	//return publications, nil
}

func (p Publication) GetPublications() ([]models.PublicationReturn, error) {
	ctx := context.Background()

	var (
		publications []models.PublicationReturn
	)

	err := p.db.Prisma.QueryRaw(`select
		P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER U on U.id = P.auth_id where u.id = ?`).Exec(ctx, &publications)

	if err != nil {
		return nil, err
	}

	return publications, nil

	//query, err := p.db.Query(`select
	//	P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.id as auth_id,u.NICK as auth_nick, u.EMAIL as auth_email
	//	from PUBLICATION as P inner join USER as U on P.auth_id = U.id`)
	//
	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(query)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//var publications []models.PublicationReturn
	//for query.Next() {
	//	var publication models.PublicationReturn
	//	if err = query.Scan(
	//		&publication.ID,
	//		&publication.Title,
	//		&publication.Description,
	//		&publication.Content,
	//		&publication.Likes,
	//		&publication.CreatedAt,
	//		&publication.AuthId,
	//		&publication.AuthNick,
	//		&publication.AuthEmail,
	//	); err != nil {
	//		return nil, err
	//	}
	//	publications = append(publications, publication)
	//}
	//return publications, nil
}

func (p Publication) UpdatePublication(id string, body models.Publication) error {
	ctx := context.Background()

	_, err := p.db.Publication.FindUnique(
		db.Publication.ID.Equals(id),
	).Update(
		db.Publication.Title.Set(body.Title),
		db.Publication.Description.Set(body.Description),
		db.Publication.Content.Set(body.Content),
	).Exec(ctx)

	if err != nil {
		return err
	}
	return nil

	//statement, err := p.db.Prepare("update PUBLICATION set content = ?, description = ?, title = ? where id = ? and auth_id = ?")
	//if err != nil {
	//	return err
	//}
	//
	//defer func(statement *sql.Stmt) {
	//	err := statement.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(statement)
	//if _, err = statement.Exec(body.Content, body.Description, body.Title, id, userId); err != nil {
	//	return err
	//}
	//
	//return nil
}

func (p Publication) GetPublication(id uint64) (models.PublicationReturn, error) {

	ctx := context.Background()

	var (
		publications models.PublicationReturn
	)

	err := p.db.Prisma.QueryRaw(`select
		P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER U on U.id = P.auth_id where u.id = ?`, id).Exec(ctx, &publications)

	if err != nil {
		return models.PublicationReturn{}, err
	}

	return publications, nil

	//query, err := p.db.Query(`select
	//	P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.id as auth_id, u.NICK as auth_nick, u.EMAIL as auth_email
	//	from PUBLICATION as P inner join USER u on u.id = P.auth_id where p.id = ?`, id)
	//
	//if err != nil {
	//	return models.PublicationReturn{}, err
	//}
	//
	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(query)
	//
	//if err != nil {
	//	return models.PublicationReturn{}, err
	//}
	//
	//var publication models.PublicationReturn
	//if query.Next() {
	//	if err = query.Scan(
	//		&publication.ID,
	//		&publication.Title,
	//		&publication.Description,
	//		&publication.Content,
	//		&publication.Likes,
	//		&publication.CreatedAt,
	//		&publication.AuthId,
	//		&publication.AuthNick,
	//		&publication.AuthEmail,
	//	); err != nil {
	//		return models.PublicationReturn{}, err
	//	}
	//}
	//
	//return publication, nil
}

func (p Publication) DeletedPublication(id string) error {

	ctx := context.Background()

	_, err := p.db.Publication.FindUnique(
		db.Publication.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
	//statement, err := p.db.Prepare("delete from PUBLICATION as P where p.id = ?")
	//
	//if err != nil {
	//	return err
	//}
	//
	//defer func(statement *sql.Stmt) {
	//	err := statement.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(statement)
	//
	//if _, err = statement.Exec(id); err != nil {
	//	return err
	//}
	//return nil
}

func (p Publication) LikePublication(id string, likes int) error {
	ctx := context.Background()

	_, err := p.db.Publication.FindUnique(
		db.Publication.ID.Equals(id),
	).Update(
		db.Publication.Like.Set(likes),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil

	//statement, err := p.db.Prepare("update PUBLICATION set `like` = ? where id = ?")
	//
	//if err != nil {
	//	return err
	//}
	//
	//defer func(statement *sql.Stmt) {
	//	err := statement.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(statement)
	//
	//if _, err = statement.Exec(likes, id); err != nil {
	//	return err
	//}
	//
	//return nil
}
