package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Reply struct {
	Message string `json:message"`
}

func main() {
	//fmt.Println("vim-go")

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	fs := http.FileServer(http.Dir("views"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	r.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		reply := Reply{
			Message: "pong",
		}

		json.NewEncoder(w).Encode(reply)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	})

	//http.ListenAndServe(":5000", r)
	http.ListenAndServe(":"+port, r)
}
