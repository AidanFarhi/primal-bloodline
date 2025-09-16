package main

import (
	"fmt"
	"log"
	"net/http"
	"primalbl/config"
	"primalbl/handler"
	"primalbl/service"
	"strconv"
)

func main() {

	conf := config.Config{}
	conf.Load()

	cs := service.NewContactService(conf)
	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))
	mux := http.NewServeMux()

	mux.Handle("/web/", fs)
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("GET /kittens", handler.KittensHandler)
	mux.HandleFunc("GET /cat-details/", handler.CatDetailsHandler)
	mux.HandleFunc("GET /inquire/", handler.InquireHandler)
	mux.HandleFunc("POST /api/contact", handler.NewContactHandler(cs))

	handlerWithRedirect := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !conf.Develop && r.Header.Get("X-Forwarded-Proto") != "https" {
			fmt.Println("redirecting")
			url := "https://" + r.Host + r.URL.RequestURI()
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		mux.ServeHTTP(w, r)
	})

	s := http.Server{
		Addr:    ":" + strconv.Itoa(conf.Port),
		Handler: handlerWithRedirect,
	}

	log.Fatal(s.ListenAndServe())
}
