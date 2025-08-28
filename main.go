package main

import (
	"fmt"
	"net/http"
	"os"
	"primalbl/config"
	"primalbl/handler"
	"primalbl/service"
)

func main() {

	conf := config.Config{}

	err := conf.Load("./config.json")

	if err != nil {
		fmt.Println("error loading config")
		os.Exit(1)
	}

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
		Addr:    ":1337",
		Handler: mux,
	}

	s.ListenAndServe()
}
