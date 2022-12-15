package main

import (
	"fmt"
	"hangman-web/hangmanClassic/HangmanStructure"
	"hangman-web/hangmanClassic/UserInput"
	"net/http"
	"strconv"
	"text/template"
)

var (
	printHidenWord string
	attempts       int
	correctLetter  bool = true
	startGame      bool = true
	endGame        bool = false
	won            bool = false
)

type gameStruct struct {
	WordToFind  string
	Attempts    int
	NumberOfPos string
	IsCorrect   bool
	BeginGame   bool
	EndGame     bool
	Won         bool
}

var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

var hangman = new(HangmanStructure.HangmanData)

func Hangman(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("html/template/hangman.html"))
	printHidenWord = ""
	for _, letter := range hangman.GetWord() {
		printHidenWord += letter
	}
	if printHidenWord == hangman.GetWordToFind() {
		endGame = true
		won = true
	}

	data := gameStruct{
		WordToFind:  printHidenWord,
		Attempts:    10 - attempts,
		NumberOfPos: strconv.Itoa(attempts),
		IsCorrect:   correctLetter,
		BeginGame:   startGame,
		EndGame:     endGame,
		Won:         won,
	}
	template.Execute(w, data)
}
func Win(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "win")

}

func Lose(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "lose")

}

func Rules(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "rules")
}

func Temp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	letter := r.FormValue("letter")

	if !UserInput.IsLetterCorrect(letter, hangman) {
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

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
