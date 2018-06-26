package main

import (
	"net/http"
	"os"
	"text/template"
)

type page struct {
	Title string
	Count int
}

func (page *page) updateCount() *page {
	page.Count++
	return page
}

func (page *page) faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/favicon.ico")
}

func (page *page) viewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page.updateCount())
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	page := page{"Go Counter", 0}

	http.HandleFunc("/", page.viewHandler)
	http.HandleFunc("/favicon.ico", page.faviconHandler) // Avoid to call page.viewHandler twice.

	http.ListenAndServe(":"+args[1], nil)
}
