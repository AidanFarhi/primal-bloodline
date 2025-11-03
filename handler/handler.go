package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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
func GetIndexPage(w http.ResponseWriter, r *http.Request) {
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

// GET /contact
func GetContactPage(w http.ResponseWriter, r *http.Request) {
	pageData := model.PageData{
		NavbarID:     alternateNavbarID,
		ImageCatName: "",
		CatName:      "",
		Cats:         []model.Kitten{},
	}
	t, _ := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/navbar.html",
		"web/templates/pages/contact.html",
	)
	t.Execute(w, pageData)
}

// GET /kittens
func GetKittensPage(cs service.CatService) http.HandlerFunc {
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
func GetCatDetailsPage(cs service.CatService) http.HandlerFunc {
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
func GetInquirePage(w http.ResponseWriter, r *http.Request) {
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
func PostContact(cs service.ContactService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		number := r.FormValue("number")
		message := r.FormValue("message")
		cs.SendMessage(name, number, message)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

// GET /api/contract
func GetContract(contractPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(contractPath)
		if err != nil {
			// TODO: serve up custom error page
			http.Error(w, "Oops! Could not get contract...", http.StatusInternalServerError)
			fmt.Println("error reading contract")
			return
		}
		defer file.Close()
		fi, err := file.Stat()
		if err != nil {
			// TODO: serve up custom error page
			http.Error(w, "Oops! Could not get file info...", http.StatusInternalServerError)
			fmt.Println("error getting file info")
			return
		}
		modTime := fi.ModTime()
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=\"PrimalBloodlineContract.pdf\"")
		http.ServeContent(w, r, "PrimalBloodlineContract.pdf", modTime, file)
	}
}
