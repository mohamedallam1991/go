package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("some.html")
	if err != nil {
		fmt.Printf("errors")
		return
	}
	t.ExecuteTemplate(w, "some.html", nil)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("starting the app")

	_ = http.ListenAndServe(":8080", nil)
	// a := Sum(5, 5)
	// fmt.Println(a)
}
