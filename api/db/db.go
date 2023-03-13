package db

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
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
	return db, err
}
