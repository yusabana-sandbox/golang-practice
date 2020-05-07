package chapter11

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func DoSamples() {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	const sql = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
}
