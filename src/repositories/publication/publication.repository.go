package publication

import (
	"api/src/models"
	"database/sql"
	"log"
)

type Publication struct {
	db *sql.DB
}

func PublicationsRepository(db *sql.DB) *Publication {
	return &Publication{
		db: db,
	}
}

func (p Publication) Create(Publication models.Publication) (uint64, error) {
	statement, err := p.db.Prepare("insert into PUBLICATION (title, auth_id, description, content, `like`) values (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	insert, err := statement.Exec(Publication.Title, Publication.AuthId, Publication.Description, Publication.Content, Publication.Likes)

	if err != nil {
		return 0, err
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(idInsert), nil
}

func (p Publication) GetPublicationByUser(id uint64) ([]models.PublicationReturn, error) {
	query, err := p.db.Query(`select
		P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER U on U.id = P.auth_id where u.id = ?`, id)

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	if err != nil {
		return nil, err
	}

	var publications []models.PublicationReturn
	for query.Next() {
		var publication models.PublicationReturn
		if err = query.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Description,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthNick,
			&publication.AuthEmail,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

func (p Publication) GetPublications() ([]models.PublicationReturn, error) {
	query, err := p.db.Query(`select
		P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER as U on P.auth_id = U.id`)

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	if err != nil {
		return nil, err
	}

	var publications []models.PublicationReturn
	for query.Next() {
		var publication models.PublicationReturn
		if err = query.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Description,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthNick,
			&publication.AuthEmail,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

func (p Publication) UpdatePublication(id, userId uint64, body models.Publication) error {
	statement, err := p.db.Prepare("update PUBLICATION set content = ?, description = ?, title = ? where id = ? and auth_id = ?")
	if err != nil {
		return err
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)
	if _, err = statement.Exec(body.Content, body.Description, body.Title, id, userId); err != nil {
		return err
	}

	return nil
}

func (p Publication) GetPublication(id, userId uint64) (models.PublicationReturn, error) {

	query, err := p.db.Query(`select
    	P.id, P.title, P.description, P.content, P.like, u.CREATED_AT, u.NICK as auth_nick, u.EMAIL as auth_email
		from PUBLICATION as P inner join USER u on u.id = P.auth_id where p.id = ?`, id)

	if err != nil {
		return models.PublicationReturn{}, err
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	if err != nil {
		return models.PublicationReturn{}, err
	}

	var publication models.PublicationReturn
	if query.Next() {
		if err = query.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Description,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthNick,
			&publication.AuthEmail,
		); err != nil {
			return models.PublicationReturn{}, err
		}
	}

	return publication, nil
}
