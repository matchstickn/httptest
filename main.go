package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn anything")
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn getting")
	})

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn getting test")
	})

	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn posting")
	})

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
