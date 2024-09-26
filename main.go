package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	// "fmt"
)
type Data struct {
	Name string
	Age int
}
// func indexHandler(w http.ResponseWriter, r *http.Request) {

// }

func ageHandler(w http.ResponseWriter, r *http.Request) {
	a := Data{Name: "Mostafa", Age: 27}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error()/*convert the error into string*/, http.StatusInternalServerError)
	}
	t.Execute(w, a)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "mostafa")
}
func  main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/age/", ageHandler)
	http.ListenAndServe(":8080", nil)


	
}