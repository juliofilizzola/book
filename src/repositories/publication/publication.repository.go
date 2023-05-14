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
