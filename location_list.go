package main

import (
	//	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getAllLocations() {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	//	var locationResponse

	//	err = json.Unmarshal(dat)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
