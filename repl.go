package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	fmt.Println("Enter one of the following commands:")

	// this loop is for you gatlin
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			comm := command.callback()
			if comm != nil {
				fmt.Println(comm)

			}
			continue
		} else {
			fmt.Println("Uknown Error")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

type commandCli struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]commandCli {
	return map[string]commandCli{

		"pokemon": {
			name:        "pokemon",
			description: "Searching for a pokemon",
			callback:    commandPokemon,
		},
		"mapf": {
			name:        "mapf",
			description: "Mapping",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exiting the pokdex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Type exit to end the pokedex ...",
			callback:    commandHelp,
		},
	}

}
