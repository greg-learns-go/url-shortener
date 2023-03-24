package urls_db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/greg-learns-go/url-shortener/shortener"
	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	Db *sql.DB
}

type Entry struct {
	ShortUrl string
	LongUrl  string
}

func CreateConnection(connString string) Connection {
	db, er := sql.Open("sqlite3", connString)
	if er != nil {
		panic(er)
	}
	return Connection{Db: db}
}

func (conn *Connection) Close() {
	conn.Db.Close()
}

func (conn *Connection) EnsureDBExists() {
	_, er := conn.Db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
		  long_url TEXT NOT NULL PRIMARY KEY,
			short_url TEXT NOT NULL
	 ) WITHOUT ROWID;

		CREATE UNIQUE INDEX IF NOT EXISTS unique_long_urls
		ON urls(long_url);

		CREATE UNIQUE INDEX IF NOT EXISTS unique_short_urls
		ON urls(short_url);
	`)
	if er != nil {
		panic(er)
	}
}

func (conn *Connection) Find(short_url string) (long_url string, er error) {
	row := conn.Db.QueryRow(
		`SELECT long_url FROM urls WHERE short_url = ? LIMIT 1`,
		short_url,
	)
	if er := row.Scan(&long_url); er != nil {
		return "", errors.New("url not found")
	}

	return
}

// Maybe later I'll come up with an interface for that?
func (conn *Connection) Insert(long_url, short_url string) (er error) {
	_, er = conn.Db.Exec(`
		INSERT INTO urls (long_url, short_url)
		VALUES (?, ?)
	`, long_url, short_url)
	return er
}

func (conn *Connection) GetAll() (results []Entry, er error) {
	var entry Entry
	rows, er := conn.Db.Query("Select * from urls")
	if er != nil {
		return nil, er
	}
	defer rows.Close()
	for rows.Next() {
		entry = Entry{}
		if er := rows.Scan(&entry.LongUrl, &entry.ShortUrl); er != nil {
			return nil, er
		}
		results = append(results, entry)
	}
	return
}

func (conn *Connection) FindOrInsert(long string) (Entry, error) {
	short := shortener.Shorten(long)

	longInDb, er := conn.Find(short)
	if er != nil {
		if er = conn.Insert(long, short); er != nil {
			fmt.Println("[ERR]", er)
		}
		return Entry{LongUrl: long, ShortUrl: short}, nil
	}

	if longInDb == long {
		// Exact duplicate, no need to do anything
		return Entry{ShortUrl: short, LongUrl: long}, nil
	}

	// conflict (long URL generated a hash that a different URL has already been saved with)
	panic("Conflict, Not implemented yet")
}
