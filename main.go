package main

import (
	"fmt"
	"log"
	"net/http"

	"pushtest/middle"
	"pushtest/rps"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/api/test", middle.Logger(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "getting test")
	}))

	router.HandleFunc("/api/{item}/", middle.Logger(func(w http.ResponseWriter, r *http.Request) {
		item := r.PathValue("item")
		outcome := rps.Play(item)
		fmt.Fprint(w, outcome)
	}))

	port := ":4000"

	fmt.Println("Listening on port, ", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
