package main

import(
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		t := template.Must(template.ParseFiles("main.html"))

		t.Execute(w, Info())

	})

	http.ListenAndServe(":8080", r)
}