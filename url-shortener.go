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
