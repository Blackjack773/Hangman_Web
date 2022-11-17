package main

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
<<<<<<< HEAD
	http.HandleFunc("/hangman", Hangman)
=======
	http.HandleFunc("/index/Game", Game)
>>>>>>> 1079b051b98ceef7d4467eafe15a11f3642d1ec9
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/template/css/"))))
	fmt.Println("http://localhost:3000 - Server started on port :3000")

	http.ListenAndServe(port, nil)
}
