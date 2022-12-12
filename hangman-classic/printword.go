package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Affichage des lettres en Ascii
// Elles sont afficher ligne par ligne pour pouvoire les afficher cote à cote
// On affiche la premiere ligne de chaque lettre puis la deuxieme ligne ...
func Prtword(tab [][]string) {
	for i := 0; i < len(tab[0]); i++ {
		for j := 0; j < len(tab); j++ {
			fmt.Printf(tab[j][i])
			// C'est pour ca qu'on a pas : fmt.Printf(tab[I][J])
		}
		fmt.Printf("\n")
	}
}

// Affichages de N lettres aléatoires pour deviner le Mot au début
func NLetter(Mot string, ascii [][]string, Affich []string, MotATrouve []string) ([][]string, []string, []string) {

	n := len(Mot)/2 - 1 // Le nombre de lettre à afficher
	tab := []int{}      // creation d'un tableau qui contiendra les indices à afficher
	boolean := false    // pour une boucle

	var Basique bool

	if len(MotATrouve) == 0 {
		Basique = false
	} else {
		Basique = true
	}

	// On affiche aucune lettre car le mot est trop petit
	if n == 0 {
		return ascii, Affich, MotATrouve
	}

	// On recupere n numéros < len(Mot)
	// Ce sont les indices des lettres à afficher
	for !boolean { // La boucle s'arrete quand tous les éléments de tab sont différents

		// recuperation au hazard des n numeros
		for i := 0; i < n; i++ {

			rand.Seed(time.Now().UnixNano()) // "planter une graine" aléatoire en fonction de l'heure
			radomInt2 := rand.Intn(len(Mot)) // random des numeros
			tab = append(tab, radomInt2)     // puis ajout a notre tableau tab
		}
		boolean = Alldiff(tab) // tous les indices doivent évidemment etre different (pour bien avoir n numeros)
	}

	// Puis on met les lettres dans le bon tableau et on ajoute les lettres au tableau affiche
	for i := 0; i < len(tab); i++ {
		if Basique {
			MotATrouve[tab[i]] = string(Mot[tab[i]]) + " "
		} else {
			ascii[tab[i]] = Lettertoascii(string(Mot[tab[i]]))
		}
		Affich = append(Affich, string(Mot[tab[i]]))
	}

	//fmt.Print(tab)
	return ascii, Affich, MotATrouve // enfin on renvoie ce dont on a besoin ...
}

// retourne true si tous les élément du tableau tab sont différents sinon retourne false
// On compare tous les éléments entre eux
func Alldiff(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] == tab[j] {
				return false
			}
		}
	}
	return true
}
