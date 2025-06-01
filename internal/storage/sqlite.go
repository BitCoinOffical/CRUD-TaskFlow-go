package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitSQLite() {
	var err error
	DB, err = sql.Open("sqlite3", "tasks.db")
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
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Cannot create table:", err)
	}
}
