package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

<<<<<<< HEAD
func Hangman(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hangman")
=======
func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Game")
>>>>>>> 1079b051b98ceef7d4467eafe15a11f3642d1ec9
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
