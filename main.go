package main

import (
	"encoding/json"
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

	router.HandleFunc("GET /api/{item}/", middle.Logger(func(w http.ResponseWriter, r *http.Request) {
		middle.CORS(w)
		w.Header().Set("Content-Type", "application/json")
		item := r.PathValue("item")
		outcome, err := rps.Play(item)
		if err != nil {
			fmt.Fprint(w, err)
			log.Println(err)
			return
		}
		json_outcome, err := json.Marshal(outcome)
		if err != nil {
			fmt.Fprint(w, err)
			log.Println(err)
			return
		}
		fmt.Fprint(w, string(json_outcome))

	}))

	port := ":4000"

	fmt.Println("Listening on port, ", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
