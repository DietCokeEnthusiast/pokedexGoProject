package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func commandPokemon() error {
	fmt.Println("Please enter a pokemon")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	pokemonName := input.Text()

	resp, err := http.Get(URL + "pokemon/" + pokemonName)

	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	var dataGotten getData

	err = json.Unmarshal(body, &dataGotten)
	fmt.Println(dataGotten, dataGotten) // dataGotten.url)
	return nil

}
