package main

import (
	"net/http"
	"primalbl/handler"
)

func main() {

	mux := http.NewServeMux()

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))

	mux.Handle("/web/", fs)

	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/kittens", handler.KittensHandler)
	mux.HandleFunc("/cat-details/", handler.CatDetailsHandler)
	mux.HandleFunc("/inquire/", handler.InquireHandler)

	s := http.Server{
		Addr:    ":1337",
		Handler: mux,
	}

	s.ListenAndServe()
}
