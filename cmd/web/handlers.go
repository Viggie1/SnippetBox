package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from SnippetBox\n"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}

	fmt.Fprintf(w, "Display a snippet with ID %d...\n", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {	
	w.Write([]byte("Displaying a form for creating a new snippet...\n"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet...\n"))
}

