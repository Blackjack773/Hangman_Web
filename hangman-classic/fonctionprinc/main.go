package main

import (
	"bufio"
	"fmt"
	"hangman"
	"io"
	"os"
	"os/exec"
	//"reflect"
)

func main() {
	for {
		str := ""
		fmt.Println("1 : Lancer une nouvelle partie")
		fmt.Println("2 : Continuer la partie sauvegarder")
		fmt.Println("3 : En savoir à propos du jeu")
		fmt.Println("4 : Quitter le jeu")
		fmt.Printf("Ecris le numéro de ce que tu veux faire : ")
		fmt.Scan(&str)
		fmt.Printf("\n")

		// Clear le terminal
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		switch str {
		case "1":

			// lancer la sauvegarde si elle est placé en parametre
			if len(os.Args) == 2 && os.Args[1] == "Save.json" {

				hangman.SaveParam("Save.json")

			} else {

				Version := true

				for Version {
					choix := ""
					fmt.Println("A quelle version du jeu voulez-vous jouer : ")
					fmt.Println("1 : Version Basique")
					fmt.Println("2 : Version Ascii")
					fmt.Println("3 : Sortir du jeu")
					fmt.Printf("Entrez votre choix : ")
					fmt.Scan(&choix)

					// Clear le terminal
					c := exec.Command("clear")
					c.Stdout = os.Stdout
					c.Run()

					// Initialisation des variables
					cpt := -1            // compteur d'erreurs
					var lettremanque int // cpt de lettres manquantes

					MotATrouver := []string{} // Mot a trouver (ce qui va etre afficher et changer)
					Ascci := [][]string{}     // Mot a trouver en Ascii-Art

					Mot := ""              // Mot choisi au hasard à deviner
					var Pendu []string     // Position du pendu (une ligne <=> une string)
					Lettre := []string{}   // pour les lettres deja dites
					var Affichage []string // tableau des n lettres afficher

					var Stop bool // bouleen qui indique si on doit sauvegarder ou non

					// Determiner le mot a deviner le mot à deviner
					Mot = hangman.Findword()
					//fmt.Println(Mot)

					// on teste que le mot soit bon
					if Mot == "Veuillez relancer le jeux" {
						fmt.Println("Veuillez relancer le jeux")
						return
					}

					// On recupere les n lettres a afficher
					n := (len(Mot) / 2) - 1
					lettremanque = len(Mot) - n // Initialisation de lettremanque

					// On recupere les positions du pendu
					Pendu = hangman.Hangmanpose()

					switch choix {
					case "1": // Version Basique

						// ajout des "_" dans le Mot a trouver
						for i := 0; i < len(Mot); i++ {
							MotATrouver = append(MotATrouver, "_ ")
						}

						// On recupere les n lettres a afficher
						Ascci, Affichage, MotATrouver = hangman.NLetter(Mot, Ascci, Affichage, MotATrouver)

						// On lance le jeu avec toutes les variables
						// Et on les recupere si jamais l'utilisateur veut sauvegarder ca partie
						cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Stop, Affichage, Ascci = hangman.Jeu(cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Affichage, Ascci)

						// Sauvegarde
						hangman.Sauveg(Stop, cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affichage, MotATrouver)

					case "2":

						// creation des "_" en Ascii
						for i := 0; i < len(Mot); i++ {
							Ascci = append(Ascci, []string{"             ", "             ", "             ", "             ", " ___________ ", "|___________|", "              "}) // 7 éléments
						}

						// On recupere les n lettres a afficher
						Ascci, Affichage, MotATrouver = hangman.NLetter(Mot, Ascci, Affichage, MotATrouver)

						// On lance le jeu
						cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Stop, Affichage, Ascci = hangman.Jeu(cpt, lettremanque, MotATrouver, Mot, Pendu, Lettre, Affichage, Ascci)

						// Sauvegarde
						hangman.Sauveg(Stop, cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affichage, MotATrouver)

					case "3":
						Version = false
					default:
						fmt.Printf("\n")
						fmt.Println(hangman.Red + "Veuillez entrer une reponse correcte svp" + hangman.Reset)
						fmt.Printf("\n")
					}
				}
			}
		case "2":
			// Lancer la sauvegarde d'une autre facon

			bouleen := true
			for bouleen {
				choix := ""
				fmt.Println("Voulez-vous continuer la partie sauvegarder ? : oui/non")
				fmt.Printf("Entrez votre choix : ")
				fmt.Scan(&choix)
				switch choix {
				case "oui":

					fmt.Println("Vous pouvez dès à présent reprendre votre partie")

					hangman.SaveParam("Save.json")

					bouleen = false

				case "non":
					bouleen = false

					// Clear le terminal
					c := exec.Command("clear")
					c.Stdout = os.Stdout
					c.Run()

				default:
					fmt.Printf("\n")
					fmt.Println(hangman.Red + "Veuillez entrer une reponse correcte svp" + hangman.Reset)
					fmt.Printf("\n")
				}
			}
		case "3":
			file, err := os.Open("readme.txt")
			if err != nil {
				panic(err)
			}

			defer file.Close()

			reader := bufio.NewReader(file)

			for {
				line, _, err := reader.ReadLine()

				if err == io.EOF {
					break
				}

				fmt.Println(string(line))
			}
		case "4":
			return

		default:
			fmt.Printf("\n")
			fmt.Println(hangman.Red + "Veuillez entrer une reponse correcte svp" + hangman.Reset)
			fmt.Printf("\n")
		}
	}
}
