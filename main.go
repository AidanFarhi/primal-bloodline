package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"primalbl/config"
	"primalbl/handler"
)

func loadConfig(configPath string, conf *config.Config) error {
	fileBytes, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBytes, conf)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	conf := config.Config{}

	err := loadConfig("./config.json", &conf)

	if err != nil {
		fmt.Println("error loading config")
		os.Exit(1)
	}

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
