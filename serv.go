package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	//Demarrage du Serveur
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomePage)
	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}
