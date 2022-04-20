package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver connection with mysql
)

// open connection with database
func Connect() (*sql.DB, error) {
	urlConn := "golang:golang@/books?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
