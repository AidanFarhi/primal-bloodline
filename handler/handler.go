package handler

import (
	"html/template"
	"net/http"
	"path"
	"primalbl/model"
	"primalbl/service"
	"strings"
)

const (
	mainNavbarID      = "main-navbar"
	alternateNavbarID = "alternate-navbar"
)

// Helper function to extract cat name from URL path
func ExtractCatNameFromPath(urlPath, prefix string) string {
	catName := strings.TrimPrefix(urlPath, prefix)
	catName = path.Clean(catName)
	return catName
}

// GET /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageData := model.PageData{
		NavbarID:     mainNavbarID,
		ImageCatName: "",
		CatName:      "",
		Cats:         []model.Kitten{},
	}
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/navbar.html",
		"web/templates/pages/index.html",
	)
	t.Execute(w, pageData)
}

// GET /kittens
func NewKittensHandler(cs service.CatService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kittens := cs.GetAllCats()
		pageData := model.PageData{
			NavbarID:     alternateNavbarID,
			ImageCatName: "",
			CatName:      "",
			Cats:         kittens,
		}
		t, _ := template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/partials/navbar.html",
			"web/templates/pages/kittens.html",
			"web/templates/partials/kitten-card.html",
		)
		t.Execute(w, pageData)
	}
}

// GET /cat-details/{catName}
func NewCatDetailsHandler(cs service.CatService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catName := ExtractCatNameFromPath(r.URL.Path, "/cat-details/")
		kittens := []model.Kitten{}
		kitten := cs.GetCatByReferenceName(catName)
		kittens = append(kittens, kitten)
		pageData := model.PageData{
			NavbarID:     alternateNavbarID,
			ImageCatName: "",
			CatName:      "",
			Cats:         kittens,
		}
		t, _ := template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/pages/kitten-details.html",
			"web/templates/partials/navbar.html",
			"web/templates/partials/kitten-detail.html",
		)
		t.Execute(w, pageData)
	}
}

// GET /inquire/{catName}
func InquireHandler(w http.ResponseWriter, r *http.Request) {
	catName := strings.TrimPrefix(r.URL.Path, "/inquire/")
	catName = path.Clean(catName)
	titleCatName := strings.ToUpper(catName[0:1]) + catName[1:]
	pageData := model.PageData{
		NavbarID:     alternateNavbarID,
		ImageCatName: catName,
		CatName:      titleCatName,
		Cats:         []model.Kitten{},
	}
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/navbar.html",
		"web/templates/pages/inquire.html",
	)
	t.Execute(w, pageData)
}

// POST /api/contact
func NewContactHandler(cs service.ContactService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		number := r.FormValue("number")
		message := r.FormValue("message")
		cs.SendMessage(name, number, message)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
