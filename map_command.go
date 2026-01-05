package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var nextURL = (URL + "location-area/")
var prevURL string

func commandMapf() error { //write function call in repl.go next pls Nathan
	fmt.Println("Going to the next page of locations")
	resp, err := http.Get(nextURL)
	if err != nil {
		log.Fatal(err)

	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	//Making a data structure to shove the json into from API
	dataGotten := getData{}
	//Must Unmarshal data to be able to use this
	err = json.Unmarshal(body, &dataGotten)

	prevURL = dataGotten.Previous
	nextURL = dataGotten.Next
	for _, loc := range dataGotten.Results {
		fmt.Println(loc.Name) // this only prints the first twenty names, make mapf(command) do next page
	}
	return nil

}
func commandMapb() error { // this will be func to show the previous page of locations
	fmt.Println("Going to previous page of locations")
	if prevURL == "" {

		fmt.Println("Youre already on the first page!")
		return nil
	}

	resp, err := http.Get(prevURL)

	if err != nil {
		log.Fatal(err)

	}
	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	//Making a data structure to shove the json into from API

	dataGotten := getData{}

	//Must Unmarshal data to be able to use this
	err = json.Unmarshal(body, &dataGotten)
	//	nextPage := dataGotten.Next

	prevURL = dataGotten.Previous
	nextURL = dataGotten.Next

	for _, loc := range dataGotten.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
