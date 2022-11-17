package Hangman_Web

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/template/css/"))))
	fmt.Println("http://localhost:3000 - Server started on port :3000")

	http.ListenAndServe(port, nil)
}
