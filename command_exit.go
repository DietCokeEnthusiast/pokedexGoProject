package main

import (
	"fmt"
	"os"
)

func commandExit() error { //write function call in repl.go next pls Nathan
	fmt.Println("Exiting the Pokedex now ...")
	os.Exit(0)
	return nil

}
