package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

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
	http.HandleFunc("/all", allLinks)
	http.ListenAndServe(":8080", nil)

	// TODO: This is not printed, check out why
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("ctrl-C to terminate")
}

func allLinks(w http.ResponseWriter, r *http.Request) {
	records, er := conn.GetAll()
	if er != nil {
		log.Fatal("Error:", er)
	}

	t := loadTemplate("all.template.html")
	t.Execute(w, records)
}

func loadTemplate(name string) (t *template.Template) {
	str, er := os.ReadFile("templates/" + name)
	if er != nil {
		panic(er)
	}
	t, er = template.New(name).Parse(string(str))
	if er != nil {
		panic(er)
	}
	return
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
			fmt.Println("[INF] POST!!!!!!", r.Form["url"][0])
			entry, er := conn.FindOrInsert(r.Form["url"][0])
			renderPostResponse(w, entry, er)
		}
	} else {
		findLinkAndServe(w, path)
	}
}

func renderPostResponse(w http.ResponseWriter, entry urls_db.Entry, er error) {
	var t *template.Template

	if er != nil {
		t = loadTemplate("submission.error.template.html")
		t.Execute(w, er)
	} else {
		t = loadTemplate("submission.success.template.html")
		t.Execute(w, entry)
	}
}

func renderSubmissionForm(w http.ResponseWriter) {
	t := loadTemplate("form.template.html")
	t.Execute(w, nil)
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
