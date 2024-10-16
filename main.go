package main

import (
	"fmt"
	"net/http"
)

func main() {
  fmt.Println("Hello, Coffee API service!")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /coffee", func(w http.ResponseWriter, r *http.Request) {	
		fmt.Fprintf(w, "this is all of the coffees!")
	})

	mux.HandleFunc("GET /coffee/{id}", func(w http.ResponseWriter, r *http.Request) {	
		id := r.PathValue("id")
		fmt.Fprintf(w, "The selected coffee (id: %s)!", id)
	})

	mux.HandleFunc("POST /coffee", func(w http.ResponseWriter, r *http.Request) {	
		fmt.Fprintf(w, "Just updating the coffees!")
	})

	// mux.HandleFunc")
	// http.ListenAndServe(":8080", mux)


	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}

