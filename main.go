package main

import "net/http"

func main() {
	
	mux := http.NewServeMux()

	fs := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))

	mux.Handle("/web/", fs)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/html/index.html")
	})

	s := http.Server{
		Addr: ":1337",
		Handler: mux,
	}

	s.ListenAndServe()
}