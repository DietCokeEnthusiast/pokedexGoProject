package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap() error { //write function call in repl.go next pls Nathan
	fmt.Println("Mapping")

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	type getData struct {
		Name      string `json:"name"`
		Height    int    `json:"height"`
		Abilities string `json:"id"`
	}

	var dataGotten getData

	err = json.Unmarshal(body, &dataGotten)
	fmt.Println(dataGotten.Name, dataGotten.Height) // dataGotten.url)
	return nil

}
