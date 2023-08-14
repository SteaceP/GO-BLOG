package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))
	tpl.ExecuteTemplate(w, "base", nil)
}
