package main

// faire un dictionnaire

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// compresser le dossier en .zip (pas en .rar (sinon modification de "\n" en "\r" ))

func Findword() string {

	// Initialisation Variables
	var Mot string
	var words []string
	Dictionnaire := [][]string{}

	// Si aucun fichier txt est placé en paramettre ou s'il y en a trop ...
	// On on creer notre propre base de mot pour que l'utilisateur puisse quand meme jouer
	if len(os.Args[1:]) != 1 {

		//on creer d'abord un dictionnaire ou on recupere chacun des mots de trois fichier.txt
		Dictionnaire = append(Dictionnaire, Readfile("words.txt"))
		Dictionnaire = append(Dictionnaire, Readfile("words2.txt"))
		Dictionnaire = append(Dictionnaire, Readfile("words3.txt"))

		// Puis on met tous les mots dans le même tableau words
		// avec deux boucle par valeur
		for _, txt := range Dictionnaire {
			for _, i := range txt {
				words = append(words, i)
			}
		}

		if len(words) != 0 { // On choisi un mot au hazard dans words

			// trirage d'un indice < len(words)
			rand.Seed(time.Now().UnixNano())
			nbr := rand.Intn(len(words))

			Mot = words[nbr]
			return Mot // puis on retourne le mot

		} else { // Si words est vide
			return "Veuillez relancer le jeux"
		}

	} else { // Sinon On prend un mot au hazard dans le fichier que l'utilisateur a entrer en paramatre

		words = Readfile(os.Args[1])

		if len(words) != 0 {

			// trirage d'un indice < len(words)
			rand.Seed(time.Now().UnixNano())
			nbr := rand.Intn(len(words))

			Mot = words[nbr]
			return Mot // puis on retourne le mot

		} else { // Si words est vide
			return "Veuillez relancer le jeux"
		}
	}
}

func Readfile(name string) []string { // name : nom d'un fichier.txt

	words := []string{} // tableau qui contiendra les mots

	contents, err := os.Open(name) // Ouverture fichier

	if err != nil { // Gestion erreur
		fmt.Println("File reading error", err)
		fmt.Println("Veuillez relencer le jeux")
		return words // on retourne les mots
	}

	defer contents.Close() // Fermeture du fichier

	scanner := bufio.NewScanner(contents) // Transformation de contents en tableau de byte
	scanner.Split(bufio.ScanWords)        // Separer le tableau de byte par ligne

	// Ajouter chaque ligne au tableau words (une ligne correspond à un mot)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil { // gestion d'erreur
		fmt.Println(err)
	}

	return words // On retourne les mots
}
