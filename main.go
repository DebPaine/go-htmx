package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct{
	Title string
	Year int
}

var films []Film

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/home", helloWorldTemplate)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getFilms(w http.ResponseWriter, r* http.Request){
	addFilms(Film{"dsfasdf", 12312})
	fmt.Fprint(w, films)
	fmt.Fprint(w, r.Method)
}

func addFilms(film Film){
	films = append(films, film)
}
func helloWorld(w http.ResponseWriter, r* http.Request){
	fmt.Fprint(w, "Hello World!\n")
	fmt.Fprint(w, r.Method)
}

func helloWorldTemplate(w http.ResponseWriter, r* http.Request){
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

