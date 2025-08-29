package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"primalbl/service"
	"strings"
)

// Helper function to extract cat name from URL path
func ExtractCatNameFromPath(urlPath, prefix string) string {
	catName := strings.TrimPrefix(urlPath, prefix)
	catName = path.Clean(catName)
	return catName
}

// GET /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/main-nav.html",
		"web/templates/index.html",
	)
	t.Execute(w, nil)
}

// GET /kittens
func KittensHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/alternate-nav.html",
		"web/templates/kittens.html",
	)
	t.Execute(w, nil)
}

// GET /cat-details/{catName}
func CatDetailsHandler(w http.ResponseWriter, r *http.Request) {
	catName := ExtractCatNameFromPath(r.URL.Path, "/cat-details/")
	templatePath := fmt.Sprintf("web/templates/%s.html", catName)
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		templatePath,
		"web/templates/partials/alternate-nav.html",
	)
	t.Execute(w, nil)
}

// GET /inquire/{catName}
func InquireHandler(w http.ResponseWriter, r *http.Request) {
	catName := strings.TrimPrefix(r.URL.Path, "/inquire/")
	catName = path.Clean(catName)
	titleCatName := strings.ToUpper(catName[0:1]) + catName[1:]
	pageData := map[string]string{
		"ImageCatName": catName,
		"CatName":      titleCatName,
	}
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/alternate-nav.html",
		"web/templates/inquire.html",
	)
	t.Execute(w, pageData)
}

// POST /api/contact
func NewContactHandler(cs service.ContactService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		cs.SendMessage(name, email, message)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
