package templates

import (
	"html/template"
)

// I decided to go with the more rigid approach, in which each template is
// a field in a struct. It's not very dynamic, but I'm playing with the
// static typing, and what benefits and trade-offs are.
type Template struct {
	All               *template.Template
	Form              *template.Template
	SubmissionError   *template.Template
	SubmissionSuccess *template.Template
}

func New(templatesLocation string) *Template {
	t := Template{}
	t.All = template.Must(template.ParseFiles(templatesLocation + "/all.template.html"))
	t.Form = template.Must(template.ParseFiles(templatesLocation + "/form.template.html"))
	t.SubmissionError = template.Must(template.ParseFiles(templatesLocation + "/submission.error.template.html"))
	t.SubmissionSuccess = template.Must(template.ParseFiles(templatesLocation + "/submission.success.template.html"))
	return &t
}
