package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPost struct {
	ID      int
	Title   string
	Content string
}

var fakePosts = []BlogPost{
	{1, "First Post", "This is the content of the first post."},
	{2, "Second Post", "This is the content of the second post."},
	// Add more fake posts here
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID := vars["id"]

	var post BlogPost
	for _, p := range fakePosts {
		if strconv.Itoa(p.ID) == blogID {
			post = p
			break
		}
	}

	tpl := template.Must(template.ParseFiles("./templates/base.html", "./templates/blog.html"))
	tpl.ExecuteTemplate(w, "base", post)
}
