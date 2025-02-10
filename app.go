package main

import (
	"net/http"
	"os"
	"primalbl/handler"

	_ "github.com/joho/godotenv/autoload"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	redirectUrl := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
}

func main() {

	fs := http.FileServer(http.Dir("web"))
	m := http.NewServeMux()

	m.Handle("/web/", http.StripPrefix("/web/", fs))

	m.HandleFunc("/", handler.HomeHandler)
	m.HandleFunc("GET /about", handler.AboutHandler)
	m.HandleFunc("GET /tree", handler.TreeHandler)

	server := http.Server{
		Addr:    "0.0.0.0:443",
		Handler: m,
	}

	go http.ListenAndServe("0.0.0.0:80", http.HandlerFunc(redirect))
	server.ListenAndServeTLS(os.Getenv("CERT_PATH"), os.Getenv("PRIVATE_KEY_PATH"))
}
