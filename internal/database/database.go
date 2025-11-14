package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Ping to verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to MariaDB!")
	return db, nil
}
