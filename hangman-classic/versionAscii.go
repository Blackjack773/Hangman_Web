package main

/*import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func JeuAscii(cpt int, lettremanque int, Ascci [][]string, Mot string, Pendu []string, Lettre []string, Affich []string, Motatr []string) (int, int, [][]string, string, []string, []string, bool, []string, []string) {

	var Stop bool = false
	var stop string = "stop"

	// Debut du jeu
	fmt.Println("Une nouvelle partie a été lancée")
	fmt.Printf("\n")
	fmt.Printf("Good luck")
	fmt.Printf("\n")

	// la boucle continue tansqu'il reste des lettres a trouver ou que le joueur n'a pas fait 10 erreurs
	for (cpt < 9) && (lettremanque != 0) {

		// Affichage du nombre d'essais restant en couleur
		if cpt >= 7 {
			fmt.Printf(Red+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		} else if cpt >= 4 && cpt < 7 {
			fmt.Printf(Yellow+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		} else {
			fmt.Printf(Green+"You have %v possible errors left\n", 10-(cpt+1))
			fmt.Println(Reset)
		}

		// Affichage du Mot en Ascii
		fmt.Printf("\n")
		Prtword(Ascci)
		fmt.Printf("\n")

		// Affichage des Lettres deja entrer par l'utilisateur
		fmt.Println("Voici les lettres déjà entrée :")
		fmt.Print(Lettre)
		//fmt.Print(Affich)
		//fmt.Print(lettremanque)
		fmt.Printf("\n")
		fmt.Printf("\n")

		// le joueur entre quelque chose
		fmt.Printf("Ecrivez une lettre ou un mot s'il vous plaît : ")
		var lettre string
		fmt.Scan(&lettre)

		// Clear le terminal
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// comparer le mot rentrer par le joueur
		if len(lettre) >= 2 {

			// comparer le stop (si le joueur veux arreter)
			if len(lettre) == 4 {

				// comparer le stop
				for i := 0; i < 4; i++ {
					if strings.ToLower(string(lettre[i])) != string(stop[i]) {
						break
					} else {
						Stop = true
						return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop, Affich, Motatr
					}
				}

				// comparer le mot de longueur 4
				var cpt2 int // compteur de lettre correspondentes
				for i := 0; i < 4; i++ {
					let := strings.ToLower(string(lettre[i]))
					if let == string(Mot[i]) {
						cpt2++
					}
				}
				if cpt2 == len(Mot) { // verification du nbr de lettre correspondentes
					lettremanque = 0
				} else {
					cpt += 2
					fmt.Println(Purple + "Ce n'était pas le bon mot" + Reset)

				}

			} else if len(lettre) == len(Mot) { // Comparer les mots de meme longeur
				var cpt2 int // compteur de lettre correspondentes
				for i := 0; i < len(Mot); i++ {
					let := strings.ToLower(string(lettre[i]))
					if let == string(Mot[i]) {
						cpt2++
					}
				}
				if cpt2 == len(Mot) { // verification du nbr de lettre correspondentes
					lettremanque = 0
				} else {
					cpt += 2
					fmt.Println(Purple + "Ce n'était pas le bon mot" + Reset)
				}
			} else if len(lettre) != len(Mot) { // Comparer les mots de longeur differentes
				cpt += 2
				fmt.Println(Purple + "Ce n'était pas le bon mot" + Reset)
			}

		} else { // Sinon comparer la lettre rentrer par le joueur

			var cpt4 int = 0 // compteur de lettre correspondentes

			lettre = strings.ToLower(lettre) // mettre la lettre en minuscule (si besoin)

			// On cherche si la lettre a déjà été rentrée
			for i := 0; i < len(Lettre); i++ {
				if lettre == Lettre[i] {
					cpt4++
					break
				}
			}

			if cpt4 == 0 { // Si elle est nouvelle

				Lettre = append(Lettre, lettre) // ajout à lettre déjà entrer à une liste

				// on transforme la lettre en Ascii dans le mot
				let := 0
				for i := 0; i < len(Mot); i++ {
					if lettre == string(Mot[i]) { // si la lettre corespond
						Ascci[i] = Lettertoascii(lettre) // transformation ascii
						lettremanque--                   // decremante lettre manquante
						let++                            // Si une lettre correspond : ++
					}
				}

				// Pour savoir si aucune lettre ne correspond et qu'il faut augmenter le cpt d'erreur de un
				if let == 0 {
					cpt++
				}

				// On regarde si la lettre était déjà afficher
				cpt5 := 0 // compter cbm de fois la lettre est afficher
				for i := 0; i < len(Affich); i++ {
					if Affich[i] == lettre {
						cpt5++
					}
				}
				lettremanque += cpt5 // Puis on ajoute a lettremanque car les lettres afficher
				//ne sont pas consider comme des lettres manquantes

			} else {
				fmt.Println(Purple + "Cette lettre à déjà été rentrée" + Reset)
			}
		}

		// Afficher le pendu avec des phrase en couleur
		if cpt == -1 {
			fmt.Printf(Green + "José se porte bien" + Reset)
			fmt.Printf("\n")
		} else {
			if cpt != 10 {
				Hang(cpt, Pendu)
			}
			fmt.Printf("\n")
			if cpt < 6 {
				fmt.Println(Yellow + "José est en danger" + Reset)
				fmt.Printf("\n")
			} else {
				fmt.Println(Red + "José est en danger" + Reset)
				fmt.Printf("\n")
			}
		}
	}

	// fin du jeu
	if cpt >= 9 { // les 10 essays ont été utilisé

		fmt.Println(Red + "Dommage, vous avez tué José. lance une nouvelle partie pour réessayer" + Reset)
		fmt.Printf("Le mot à trouver était : %v \n", Mot)

		// Affichage du mot complet en Ascii-Art
		for i := 0; i < len(Mot); i++ {
			Ascci[i] = Lettertoascii(string(Mot[i]))
		}
		Prtword(Ascci)
		fmt.Printf("\n")
		fmt.Printf("\n")

		return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop, Affich, Motatr // return

	} else if lettremanque == 0 { // Toutes les lettres ont été trouvée

		// Affichage du Mot complet en Ascci-Art
		for i := 0; i < len(Mot); i++ {
			Ascci[i] = Lettertoascii(string(Mot[i]))
		}
		Prtword(Ascci)

		fmt.Println(Green + "Bravo, Tu as sauvé José" + Reset)
		fmt.Printf("\n")
		fmt.Printf("Tu as trouvé le bon mot qui était %v", Mot)
		fmt.Printf("\n")
		fmt.Printf("\n")
		return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop, Affich, Motatr
	}
	return cpt, lettremanque, Ascci, Mot, Pendu, Lettre, Stop, Affich, Motatr
}*/
