package main

import (
	"fmt"
	"log"
	"net/http"
	"pushtest/rps"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "anything")
	})

	router.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "getting")
	})

	router.HandleFunc("GET /api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "getting test")
	})

	router.HandleFunc("POST /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "posting")
	})

	router.HandleFunc("/api/{item}/", func(w http.ResponseWriter, r *http.Request) {
		item := r.PathValue("item")
		outcome := rps.Play(item)
		fmt.Fprint(w, outcome)
	})

	port := ":4000"

	fmt.Println("Listening on port, ", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
