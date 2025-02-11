package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func NewHandler(pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(
			"web/templates/header.html",
			"web/templates/menu.html",
			"web/templates/base.html",
			fmt.Sprintf("web/templates/%s.html", pageName),
		)
		t.ExecuteTemplate(w, "base", nil)
	}
}
