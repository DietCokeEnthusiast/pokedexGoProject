package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandPokemon() error {
	fmt.Println("Please enter a pokemon")

	input = bufio.NewScanner()

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	type getData struct {
		Name   string `json:"name"`
		Height int    `json:"height"`
	}

	var dataGotten getData

	err = json.Unmarshal(body, &dataGotten)
	fmt.Println(dataGotten.Name, dataGotten.Height) // dataGotten.url)
	return nil

}
