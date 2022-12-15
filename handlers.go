package main

import (
	"fmt"
	"hangman-web/hangmanClassic/HangmanStructure"
	"hangman-web/hangmanClassic/UserInput"
	"net/http"
	"strconv"
	"text/template"
)

var ( // variables du jeu
	printHidenWord string
	attempts       int
	correctLetter  bool = true
	startGame      bool = true
	endGame        bool = false
	won            bool = false
	hangman             = new(HangmanStructure.HangmanData)
)

type gameStruct struct { // structure des paramètres de jeu
	WordToFind  string
	Attempts    int
	NumberOfPos string
	IsCorrect   bool
	BeginGame   bool
	EndGame     bool
	Won         bool
}

var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'accueil
	renderTemplate(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'Info
	renderTemplate(w, "about")
}

func Hangman(w http.ResponseWriter, r *http.Request) { //renvoies à la page de Jeu

	template := template.Must(template.ParseFiles("html/template/hangman.html"))
	printHidenWord = ""
	for _, letter := range hangman.GetWord() { // ajout de lettre
		printHidenWord += letter
	}
	if printHidenWord == hangman.GetWordToFind() { // Fin de la partie, victoire ou défaite
		endGame = true
		won = true
	}

	data := gameStruct{ // la variable data prend en compte la struct de jeu
		WordToFind:  printHidenWord,         // mot à deviner
		Attempts:    10 - attempts,          //nb d'essais
		NumberOfPos: strconv.Itoa(attempts), //  numero  pos image pendu
		IsCorrect:   correctLetter,          // est-ce que la lettre devinée est correcte
		BeginGame:   startGame,              // Commencement de la partie
		EndGame:     endGame,                // Fin de la partie
		Won:         won,                    // Victoire ou Défaite
	}
	template.Execute(w, data)
}
func Win(w http.ResponseWriter, r *http.Request) { // renvoies à la page de victoire
	renderTemplate(w, "win")

}

func Lose(w http.ResponseWriter, r *http.Request) { // renvoies à la page de défaite
	renderTemplate(w, "lose")

}

func Rules(w http.ResponseWriter, r *http.Request) { // renvoies à la page de règles
	renderTemplate(w, "rules")
}

func Temp(w http.ResponseWriter, r *http.Request) { // page de chargement temporaire
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	letter := r.FormValue("letter") // récupération input

	if !UserInput.IsLetterCorrect(letter, hangman) { // vérification état de jeu
		attempts += 1
		correctLetter = false
		startGame = false
		if attempts == 10 {
			endGame = true
		}
	} else {
		correctLetter = true
		startGame = false
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string) { // Parser fichie
	t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
