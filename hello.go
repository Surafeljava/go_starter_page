package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/name", name)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/starter", starter)

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe("127.0.0.1:8080", mux)
}

var templ = template.Must(template.ParseFiles("temp.html"))
var templ2 = template.Must(template.ParseFiles("starter.html"))

func login(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

func starter(w http.ResponseWriter, r *http.Request) {
	templ2.Execute(w, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World Page</h1> <p> Hello world from go-lang ... welcome!! </p>"))
}

func name(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Surafel Kindu</h1> <p> Some description about me!!</p>"))
}
