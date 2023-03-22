package main

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var conn Connection = CreateConnection("file:database.db")

func init() {
	conn.EnsureDBExists()
}

func main() {
	defer conn.Close()

	fmt.Println(conn.GetAll())

	url, er := conn.Find("so")
	if er != nil {
		fmt.Println("Find:", er)
	}
	fmt.Println("Find:", url)

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)

	// TODO: This is not printed, check out why
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("ctrl-C to terminate")
}

func sroot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var response string
		url, er := conn.Find(strings.Trim(
			r.URL.EscapedPath(), "/",
		))
		if er != nil {
			response = "Can't find this shortened URL"
		} else {
			response = "it looks like you're looking for " + url
		}
		w.Write([]byte(response))
	}
}
