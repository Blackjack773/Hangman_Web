package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//"io"

//"devt.de/krotik/common/termutil/getch"

/*func hangman(hp int) {
	f, _ := os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var line int
	a := 71 - hp*8
	b := 78 - hp*8
	for scanner.Scan() {
		if line >= a && line <= b {
			fmt.Println(scanner.Text())
		}
		line++
	}
}*/

/*func Clear() {
	var KeyPress *getch.KeyEvent

	if err := getch.Start(); err != nil {
		fmt.Println(err)
		return
	}
	defer getch.Stop()

	KeyPress, _ = getch.Getch()
	fmt.Print(KeyPress)
}*/

/*
	func Hang(hp int, pose []string) {
		a := 0 + hp*8
		b := 7 + hp*8
		for i := a; i <= b; i++ {
			fmt.Print(pose[i])
			fmt.Printf("\n")
		}

}
*/
/*type Save struct {
	cptt            int
	lettremanquante int
	MotATrouve      []string
	Asccii          [][]string
	Mots            string
	Pend            []string
	Lett            []string
}

func Unmarshal() {
	empJsonData := `{"Name":"George Smith","Age":30,}`
	empBytes := []byte(empJsonData)
	var emp Save
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
}

func Marshal(emp Save) []byte {
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(empData))
	return empData
}*/
type Identite struct {
	Name string
	Age  int
}

func Unmarshal(data []byte) Identite {
	var emp Identite
	json.Unmarshal(data, &emp)
	return emp
}

func Marshal(emp Identite) []byte {
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return empData
}

func main() {

	file, err := os.Create("Save.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File created successfully")

	defer file.Close()

	emp := Identite{Name: "George Smith", Age: 30} // creation d'une nouvelle structure

	marshaled_data := Marshal(emp) // transformation en byte

	read, err := os.OpenFile("Save.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //Open fichier; s il existe pas : il est creer

	defer read.Close() // on ferme automatiquement à la fin de notre programme

	if err != nil {
		panic(err)
	}

	// Ecrire dans le fichier text
	_, err = file.WriteString(string(marshaled_data))
	if err != nil {
		panic(err)
	}

	// Lire le fichier txt
	data, err := ioutil.ReadFile("Save.txt") // lire le fichier
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
	fmt.Println(string(marshaled_data)) // conversion de byte en string

	unmarshdata := Unmarshal(data)
	fmt.Print(unmarshdata.Age)
	fmt.Print(unmarshdata.Name)
}

//func main() {

//pose := []string{}
/*pose = hangman.Hangmanpose()
fmt.Print(pose)
for i := 0; i < len(pose); i++ {
	fmt.Print(pose[i])
	fmt.Printf("\n")
}*/

/*file, err := os.Open("hangman.txt")
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
	pose = append(pose, string(line))
	//fmt.Printf("%s \n", line)
}

for i := 0; i < 10; i++ {
	hang(i, pose)
}

c := exec.Command("clear")
c.Stdout = os.Stdout
c.Run()*/

// Creation des niveux d'erreurs
/*content, err := os.Open("hangman.txt")
if err != nil {
	return
}
defer content.Close()

scanner := bufio.NewScanner(content)
//scanner.Split(bufio.ScanWords)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	fmt.Println(err)
}*/
//pendu := hangman.Pendu(content)

/*f, err := os.Open("words.txt")

if err != nil {
	fmt.Println(err)
}

defer f.Close()

scanner := bufio.NewScanner(f)
scanner.Split(bufio.ScanWords)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	fmt.Println(err)
}*/

/*nbr := []int{}
for i := 0; i < len(nbr); i++ {
	fmt.Print(nbr[i])
}*/

/*Motdev := []string{}

if len(os.Args[1:]) == 0 {
	fmt.Printf("File name missing\n")
} else if len(os.Args[1:]) > 1 {
	fmt.Printf("Too many arguments\n")
} else {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	Mot := hangman.Findword(contents)
	fmt.Printf("Le mot à trouver est : %v", Mot)
	fmt.Printf("\n")

	for i := 0; i < len(Mot); i++ {
		//Ascci = append(Ascci, []string{})
		Motdev = append(Motdev, "_")
	}

	lettremanque := 0
	// affichage des n lettre Ascii
	n := (len(Mot) / 2) - 1
	fmt.Print((len(Mot) / 2) - 1)

	fmt.Printf("Les n lettres à afficher : %v", n)
	fmt.Printf("\n")

	lettremanque = len(Mot) - n
	fmt.Printf("Lettres manquante : %v", lettremanque)
	fmt.Printf("\n")

	Motdev = hangman.NLetter(Mot, Motdev)
	fmt.Print(Motdev)
	fmt.Printf("\n")


}*/

//hangman := ReadFile(hangman.txt)
//tab := [][]string{[]string{" _ ", "| |", "| |", "| |", "|_|", "(_)"}, []string{"   _  _    ", " _| || |_  ", "|_  __  _| ", " _| || |_  ", "|_  __  _| ", "  |_||_|   "}}
//Ascci := [][]string{}
/*for i := 97; i <= 112; i++ {
      Ascci = append(Ascci, hangman.Lettertoascii(string(rune(i))))
  }
*/
//fmt.Print(Ascci)

/*var e int
for w := "ab"; e < len(w); {
	Ascci = append(Ascci, hangman.Lettertoascii(string(rune(w[e]))))
	e++
}
for i := 0; i < len(Ascci[0]); i++ {
	for j := 0; j < len(Ascci); j++ {
		fmt.Printf(Ascci[j][i])
		//fmt.Printf("@")
	}
	fmt.Printf("\n")
}*/

/*fmt.Println("Merci d'entrer un caractère à convertir :")
  letter := ""
  fmt.Scan(&letter)
  tab := hangman.Lettertoascii(letter)
  fmt.Print(tab)*/

//}
