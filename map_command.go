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

	resp, err := http.Get(URL + "location/")

	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	type getData struct {
		Name string `json:"name"`
	}

	var dataGotten getData

	err = json.Unmarshal(body, &dataGotten)
	fmt.Println(dataGotten.Name)
	return nil

}
