package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	fmt.Println("HELELo")
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
			description: "Searching for pokemon",
			callback:    commandPokemon,
		},
		"map": {
			name:        "map",
			description: "Mapping",
			callback:    commandMap,
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
