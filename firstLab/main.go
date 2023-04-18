package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from server!")
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	nameForm := r.FormValue("name")
	address := r.FormValue("addr")

	myVar := map[string]interface{}{"name": nameForm, "addr": address}

	outputHTML(w, "./static/formComplete.html", myVar)
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)

	if err != nil {
		log.Fatal(err)
	}

	errExecute := t.Execute(w, data)

	if errExecute != nil {
		log.Fatal(err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/form", formHandle)
	http.ListenAndServe(":8080", nil)
}
