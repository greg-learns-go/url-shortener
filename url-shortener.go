package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/greg-learns-go/url-shortener/shortener"
	"github.com/greg-learns-go/url-shortener/templates"
	"github.com/greg-learns-go/url-shortener/urls_db"
)

var conn urls_db.Connection = urls_db.CreateConnection("file:database.db")
var t = templates.New("./templates/") // TODO: make it not global

func init() {
	conn.EnsureDBExists()
}

func main() {
	defer conn.Close()

	fmt.Println(conn.GetAll())

	http.HandleFunc("/", sroot)
	http.HandleFunc("/all", allLinks)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("ctrl-C to terminate")

	// TODO: explore graceful shutdown https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve/42533360#42533360
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func allLinks(w http.ResponseWriter, r *http.Request) {
	records, er := conn.GetAll()
	if er != nil {
		log.Fatal("Error:", er)
	}

	t.All.Execute(w, records)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.EscapedPath()

	fmt.Println("EscapedPath: ", path)

	if path == "/" {
		fmt.Println("[INF] / serve form or accept POST submission")
		switch r.Method {
		case "GET":
			renderSubmissionForm(w)
		case "POST":
			if er := r.ParseForm(); er != nil {
				fmt.Println(er)
			}
			fmt.Println("[INF] POST /", r.Form["url"][0])
			entry, er := conn.FindOrInsert(r.Form["url"][0], shortener.New())
			renderPostResponse(w, entry, er)
		}
	} else {
		findLinkAndServe(w, path)
	}
}

func renderPostResponse(w http.ResponseWriter, entry urls_db.Entry, er error) {
	if er != nil {
		t.SubmissionError.Execute(w, er)
	} else {
		t.SubmissionSuccess.Execute(w, entry)
	}
}

func renderSubmissionForm(w http.ResponseWriter) {
	t.Form.Execute(w, nil)
}

func findLinkAndServe(w http.ResponseWriter, path string) {
	url, er := conn.Find(strings.Trim(path, "/"))
	var response string
	if er != nil {
		response = "Can't find this shortened URL"
	} else {
		response = "it looks like you're looking for " + url
	}
	w.Write([]byte(response))
}
