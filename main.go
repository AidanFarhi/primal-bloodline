package main

import (
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

	mux := http.NewServeMux()

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))

	mux.Handle("/web/", fs)

	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("GET /kittens", handler.KittensHandler)
	mux.HandleFunc("GET /cat-details/", handler.CatDetailsHandler)
	mux.HandleFunc("GET /inquire/", handler.InquireHandler)
	mux.HandleFunc("POST /api/contact", handler.NewContactHandler(cs))

	s := http.Server{
		Addr:    ":" + strconv.Itoa(conf.Port),
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
