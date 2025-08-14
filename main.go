package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/main-nav.html",
		"web/templates/index.html",
	)
	t.Execute(w, nil)
}

func KittensHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/alternate-nav.html",
		"web/templates/kittens.html",
	)
	t.Execute(w, nil)
}

func CatDetailsHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/cat-details/")
	name = path.Clean(name)
	templatePath := fmt.Sprintf("web/templates/%s.html", name)
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		templatePath,
		"web/templates/partials/alternate-nav.html",
	)
	t.Execute(w, nil)
}

func InquireHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/alternate-nav.html",
		"web/templates/inquire.html",
	)
	t.Execute(w, nil)
}

func main() {

	mux := http.NewServeMux()

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))

	mux.Handle("/web/", fs)

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/kittens", KittensHandler)
	mux.HandleFunc("/cat-details/", CatDetailsHandler)
	mux.HandleFunc("/inquire", InquireHandler)

	s := http.Server{
		Addr:    ":1337",
		Handler: mux,
	}

	s.ListenAndServe()
}
