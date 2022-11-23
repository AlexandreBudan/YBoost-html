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
	http.HandleFunc("/Sneakers", Sneakers)
	http.HandleFunc("/Teams", Teams)
	http.HandleFunc("/Players", Players)
	http.HandleFunc("/About", About)
	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}

func Sneakers(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}

func Teams(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}

func Players(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}

func About(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./HomePage.html")

	tmpl.Execute(w, "")
}
