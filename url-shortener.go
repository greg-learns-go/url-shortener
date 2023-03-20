package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func init() {
	database = EnsureDBExists()
}

func main() {
	defer database.Close()

	fmt.Println(GetAll(database))

	fmt.Println("Test")
}
