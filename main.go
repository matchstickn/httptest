package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn anything")
	})

	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn getting")
	})

	mux.HandleFunc("GET /api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn getting test")
	})

	mux.HandleFunc("POST /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "stkn posting")
	})

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
