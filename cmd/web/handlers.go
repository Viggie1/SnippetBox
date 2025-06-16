package main

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"
)

	func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	//we use the ... to pass the contents
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}

	fmt.Fprintf(w, "Display a snippet with ID %d...\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {	
	w.Write([]byte("Displaying a form for creating a new snippet...\n"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet...\n"))
}

