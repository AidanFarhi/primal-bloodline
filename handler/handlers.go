package handler

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/header.html",
		"web/templates/base.html",
		"web/templates/home.html",
	)
	t.ExecuteTemplate(w, "base", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/header.html",
		"web/templates/base.html",
		"web/templates/about.html",
	)
	t.ExecuteTemplate(w, "base", nil)
}
