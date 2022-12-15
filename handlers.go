package main

import (
	"fmt"
	"hangman-web/hangmanClassic/HangmanStructure"
	"hangman-web/hangmanClassic/UserInput"
	"net/http"
	"strconv"
	"text/template"
)

<<<<<<< Updated upstream
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
=======
var ( // variables du jeu
	newString string 
	attempts  int
	isCorrect bool = true
	beginGame bool = true
	endGame   bool = false
	won       bool = false
)

type indexPageData struct { // structure des paramètres de jeu
	WordToFind string
	Attempts   int
	ImageName  string
	IsCorrect  bool
	BeginGame  bool
	EndGame    bool
	Won        bool
>>>>>>> Stashed changes
}

var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'accueil
	renderTemplate(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'Info
	renderTemplate(w, "about")
}

var hangman = new(HangmanStructure.HangmanData) // initialisation des variables du jeu

func Hangman(w http.ResponseWriter, r *http.Request) { //renvoies à la page de Jeu

	template := template.Must(template.ParseFiles("html/template/hangman.html"))
<<<<<<< Updated upstream
	printHidenWord = ""
	for _, letter := range hangman.GetWord() {
		printHidenWord += letter
	}
	if printHidenWord == hangman.GetWordToFind() {
=======
	newString = ""
	for _, letter := range hangman.GetWord() { // ajout de lettre 
		newString += letter
	}
	if newString == hangman.GetWordToFind() { // Fin de la partie, victoire ou défaite
>>>>>>> Stashed changes
		endGame = true
		won = true
	}

<<<<<<< Updated upstream
	data := gameStruct{
		WordToFind:  printHidenWord,
		Attempts:    10 - attempts,
		NumberOfPos: strconv.Itoa(attempts),
		IsCorrect:   correctLetter,
		BeginGame:   startGame,
		EndGame:     endGame,
		Won:         won,
=======
	data := indexPageData{ // la variable data prend en compte la struct de jeu
		WordToFind: newString, // mot à deviner
		Attempts:   10 - attempts, //nb d'essais
		ImageName:  strconv.Itoa(attempts), // image du pendu
		IsCorrect:  isCorrect, // est-ce que la lettre devinée est correcte
		BeginGame:  beginGame, // Commencement de la partie
		EndGame:    endGame, // Fin de la partie
		Won:        won, // Victoire ou Défaite
>>>>>>> Stashed changes
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

func renderTemplate(w http.ResponseWriter, tmpl string) { // Parser fichier
	t, err := template.ParseFiles("./html/template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

