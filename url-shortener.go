package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func init() {
	d, er := sql.Open("sqlite3", "file:database.db")
	if er != nil {
		panic(er)
	}
	database = d

	database.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
		  long_url TEXT NOT NULL PRIMARY KEY,
			short_url TEXT NOT NULL
	 ) WITHOUT ROWID;
	`)

	database.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS unique_long_urls
		ON urls(long_url)
	`)

	database.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS unique_short_urls
		ON urls(short_url)
	`)
}

func main() {
	defer database.Close()

	rows, er := database.Query("Select * from urls")
	if er != nil {
		panic(er)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			long_url  string
			short_url string
		)
		if er := rows.Scan(&long_url, &short_url); er != nil {
			panic(er)
		}
		fmt.Println(long_url, short_url)

	}
	// fmt.Println("Rows:", rows)

	fmt.Println("Test")
}
