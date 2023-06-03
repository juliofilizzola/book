package internal

import (
	"api/cmd/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ClientDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.UrlDatabase)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	log.Println("Successfully connected to database!")
	return db, err
}
