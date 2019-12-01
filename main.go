package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
