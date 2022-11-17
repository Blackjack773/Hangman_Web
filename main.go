package Hangman_Web

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("http://localhost:3000 - Server started on port :3000")

	http.ListenAndServe(port, nil)
}
