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

func Hangman(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hangman")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}