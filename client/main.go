package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Year      int8   `json:"year,omitempty"`
}

func main() {
	body, err := http.Get("http://backinsano.webnetwork.com.br:3005")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if body.StatusCode != 200 {
		fmt.Println("No success", body.StatusCode)
		return
	}

	jsonclient, err := io.ReadAll(body.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var users []User

	err = json.Unmarshal(jsonclient, &users)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(users)

}
