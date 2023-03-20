package main

import (
	"database/sql"
	"errors"
)

type Entry struct {
	short_url string
	long_url  string
}

func EnsureDBExists() *sql.DB {
	db, er := sql.Open("sqlite3", "file:database.db")
	if er != nil {
		panic(er)
	}

	db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
		  long_url TEXT NOT NULL PRIMARY KEY,
			short_url TEXT NOT NULL
	 ) WITHOUT ROWID;
	`)

	db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS unique_long_urls
		ON urls(long_url)
	`)

	db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS unique_short_urls
		ON urls(short_url)
	`)

	return db
}

func Find(db *sql.DB, short_url string) (long_url string, er error) {
	row := db.QueryRow(
		`SELECT long_url FROM urls WHERE short_url = ? LIMIT 1`,
		short_url,
	)
	if er := row.Scan(&long_url); er != nil {
		return "", errors.New("url not found")
	}

	return
}

// Maybe later I'll come up with an interface for that?
func Insert(db *sql.DB, long_url, short_url string) (er error) {
	_, er = db.Exec(`
		INSERT INTO urls (long_url, short_url)
		VALUES (?, ?)
	`, long_url, short_url)
	return er
}

func GetAll(db *sql.DB) (results []Entry, er error) {
	rows, er := database.Query("Select * from urls")
	if er != nil {
		return nil, er
	}
	defer rows.Close()
	for rows.Next() {
		var (
			long_url  string
			short_url string
		)
		if er := rows.Scan(&long_url, &short_url); er != nil {
			return nil, er
		}
		results = append(results, Entry{long_url: long_url, short_url: short_url})
	}
	return
}
