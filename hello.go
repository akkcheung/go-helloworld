package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//fmt.Println("vim-go")

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.ListenAndServe(":5000", r)
}
