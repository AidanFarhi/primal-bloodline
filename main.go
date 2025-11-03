package main

import (
	"log"
	"net/http"
	"primalbl/config"
	"primalbl/handler"
	"primalbl/repo"
	"primalbl/service"
	"strconv"
)

func main() {

	conf := config.Config{}
	conf.Load()

	contactService := service.NewContactService(conf)
	catRepo := repo.NewCatRepository(conf)
	catService := service.NewCatService(catRepo)

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))
	mux := http.NewServeMux()

	mux.Handle("/web/", fs)
	mux.HandleFunc("/", handler.GetIndexPage)
	mux.HandleFunc("GET /kittens", handler.GetKittensPage(catService))
	mux.HandleFunc("GET /cat-details/", handler.GetCatDetailsPage(catService))
	mux.HandleFunc("GET /inquire/", handler.GetInquirePage)
	mux.HandleFunc("GET /contact", handler.GetContactPage)
	mux.HandleFunc("POST /api/contact", handler.PostContact(contactService))
	mux.HandleFunc("GET /api/contract", handler.GetContract(conf.ContractPath))

	handlerWithRedirect := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !conf.Develop && r.Header.Get("X-Forwarded-Proto") != "https" {
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
