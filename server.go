package main

import (
	"net/http"
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

}

func (page *page) viewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page.updateCount())
	if err != nil {
		panic(err)
	}
}

func main() {
	page := page{"Http Server", 0}

	http.HandleFunc("/", page.viewHandler)
	http.HandleFunc("/favicon.ico", page.faviconHandler) // Avoid to call page.viewHandler twice.

	http.ListenAndServe(":8080", nil)
}
