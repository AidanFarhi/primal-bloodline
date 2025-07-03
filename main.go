package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/templates/layout.html", "web/templates/index.html")
	t.Execute(w, nil)
}

func main() {
	
	mux := http.NewServeMux()

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))

	mux.Handle("/web/", fs)

	mux.HandleFunc("/", IndexHandler)

	s := http.Server{
		Addr: ":1337",
		Handler: mux,
	}

	s.ListenAndServe()
}