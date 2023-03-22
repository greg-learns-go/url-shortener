package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/greg-learns-go/url-shortener/urls_db"
)

var conn urls_db.Connection = urls_db.CreateConnection("file:database.db")

func init() {
	conn.EnsureDBExists()
}

func main() {
	defer conn.Close()

	fmt.Println(conn.GetAll())

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)

	// TODO: This is not printed, check out why
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("ctrl-C to terminate")
}

func sroot(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// When POST, expect long URL, present back short URL
	//   -- maybe implement own hashing function, that would
	//   -- differently calculate each char of short_url (go routines?)
	//   -- save current length, when options are exhausted (too many collisions)
	//   -- calculate has for short URL with one more character
	// When GET - respond with 30X redirect
	// When GET ?all - don't look for short url, but present all known URLs

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
