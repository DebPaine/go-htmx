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

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/home", helloWorldTemplate)
	http.HandleFunc("/films", getFilms)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getFilms(w http.ResponseWriter, r* http.Request){
	films := map[string][]Film{
		"films": {
			{Title: "Godzilla vs Kong", Year: 2024},
			{Title: "The Batman", Year: 2022},
		},
	}

	templ, err := template.ParseFiles("films.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err = templ.Execute(w, films)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

