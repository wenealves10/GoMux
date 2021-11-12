package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Year      int8   `json:"year,omitempty"`
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "application/json")

	json.NewEncoder(w).Encode([]User{{
		ID:        "15224",
		Firstname: "Wene",
		Lastname:  "Alves",
		Year:      19,
	}, {
		ID:        "1588",
		Firstname: "Ismael",
		Lastname:  "Albert",
		Year:      18,
	}})

}

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", HandlerUser).Methods("GET")
	log.Println("Server is running..")
	log.Fatal(http.ListenAndServe(":3005", muxRouter))
}
