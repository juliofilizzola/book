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
    	P.id, P.title, P.auth_id, P.description, P.content, P.'like', P.created_at, U.id as AuthId, U.name as AuthName, U.nick as AuthNick, U.email as AuthEmail
		from PUBLICATION as P inner join USER U on P.auth_id = U.id where auth_id = ?`, id)

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
			&publication.Content,
			&publication.Description,
			&publication.AuthId,
			&publication.AuthName,
			&publication.AuthEmail,
			&publication.AuthNick,
			&publication.CreatedAt,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

func (p Publication) GetPublication() ([]models.PublicationReturn, error) {
	query, err := p.db.Query(`select
    	P.id, P.title, P.auth_id, P.description, P.content, P.'like', P.created_at, U.id as AuthId, U.name as AuthName, U.nick as AuthNick, U.email as AuthEmail
		from PUBLICATION as P inner join USER U on P.auth_id = U.id`)

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
			&publication.Content,
			&publication.Description,
			&publication.AuthId,
			&publication.AuthName,
			&publication.AuthEmail,
			&publication.AuthNick,
			&publication.CreatedAt,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}

//func (p Publication) UpdatePublication() error {
//
//}
