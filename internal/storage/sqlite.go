package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	DB *sql.DB
}

func InitSQLite(path string) *DataBase {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Cannot open DB:", err)
	}
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description VARCHAR,
		priority VARCHAR,
		status VARCHAR,
		title VARCHAR NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Cannot create table:", err)
	}
	return &DataBase{DB: db}
}
