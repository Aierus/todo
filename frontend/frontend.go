package frontend

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreateForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Title:", r.Form["Title"])
		fmt.Println("Description:", r.Form["Description"])
	}
}
