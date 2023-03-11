package db

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.URL_DATABASE)
	if error != nil {
		return nil, error
	}
	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}
	return db, error
}
