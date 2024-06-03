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

	home_page := http.FileServer(http.Dir("./"))

	router.Handle("/", home_page)
	router.HandleFunc("GET /api/rps/{item}/", middle.Logger(func(w http.ResponseWriter, r *http.Request) {
		middle.CORS(w)
		w.Header().Set("Content-Type", "application/json")
		item := r.PathValue("item")
		if item == "hi" {
			fmt.Fprint(w, "Rock Paper Scissors api \nReplace /hi to paths /rock, /paper, /scissors to choose your input \nThe bot automatically chooses a random input")
			return
		}
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
