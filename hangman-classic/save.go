package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Creation d'une structure pour sauvegarder les donner utile pour jouer
type Save struct {
	Cptt            int
	Lettremanquante int
	MotATrouve      []string
	Asccii          [][]string
	Mots            string
	Pend            []string
	Lett            []string
	Affiche         []string
}

// Detransformation de notre objet pour pouvoir relancer le jeu avec les données sauvegarder
func Detransformation(data []byte) Save {
	var emp Save
	json.Unmarshal(data, &emp)
	return emp
}

// Transformation des de l'objet en JSON string pour pouvoir l'ecrire dans un fichier.txt
func Transformation(emp Save) []byte {
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return empData
}

func Sauveg(Stop bool, cpt int, lettremanque int, Ascci [][]string, Mot string, Pendu []string, Lettre []string, Affich []string, Motatr []string) {
	if Stop {
		boolean := true
		for boolean {
			// Demande s'il veut sauvegarder
			fmt.Println("Voulez vous sauvegarder votre partie : oui / non ")
			choice := ""
			fmt.Scan(&choice)

			switch choice {
			case "oui":

				// fonction sauvegarde
				Sauvegarder(cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Affich, Motatr)

				fmt.Println("Votre Partie a été sauvegarder")

				boolean = false
			case "non":
				fmt.Println("Votre partie n'a pas été sauvegarder")
				boolean = false
			default:
				fmt.Printf("\n")
				fmt.Println(Red + "Veuillez entrer une reponse correcte svp" + Reset)
				fmt.Printf("\n")
			}
		}
	}
}

func Sauvegarder(cpt int, lettremanque int, Ascci [][]string, Mot string, Pendu []string, Lettre []string, Affich []string, Motatr []string) {
	file, err := os.Create("Save.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sauvegarde Save
	sauvegarde.Cptt = cpt
	sauvegarde.Lettremanquante = lettremanque
	sauvegarde.MotATrouve = Motatr
	sauvegarde.Asccii = Ascci
	sauvegarde.Lett = Lettre
	sauvegarde.Mots = Mot
	sauvegarde.Pend = Pendu
	sauvegarde.Affiche = Affich

	marshaled_data := Transformation(sauvegarde) // transformation en byte

	read, err := os.OpenFile("Save.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //Open fichier; s il existe pas : il est creer

	defer read.Close() // on ferme automatiquement à la fin de notre programme

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(marshaled_data))
	if err != nil {
		panic(err)
	}
}

func SaveParam(name string) {
	data, err := ioutil.ReadFile(name) // lire le fichier
	if err != nil {
		fmt.Println(err)
	}

	unmarshsave := Detransformation(data) // On recupere L'objet écrit dans le fichier.txt

	// Puis on lance la version qui correspond : Ascii ou Basique
	if len(unmarshsave.Asccii) == 0 {
		Jeu(unmarshsave.Cptt, unmarshsave.Lettremanquante, unmarshsave.MotATrouve, unmarshsave.Mots, unmarshsave.Pend, unmarshsave.Lett, unmarshsave.Affiche, unmarshsave.Asccii)
	} else {
		Jeu(unmarshsave.Cptt, unmarshsave.Lettremanquante, unmarshsave.MotATrouve, unmarshsave.Mots, unmarshsave.Pend, unmarshsave.Lett, unmarshsave.Affiche, unmarshsave.Asccii)
	}
}
