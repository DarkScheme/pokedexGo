package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	newMap := map[string]cliCommand{
					"exit": {
					name: "exit",
					description: "Exit the Pokedex",
					callback: commandExit,
					},
					"help": {
						name: "help",
						description: "Displays a help message",
						callback: commandHelp,
					},
				}
	return newMap
}


func startRepl() {
	// creates the map
	



	// creates a buffer scanner
	scanner := bufio.NewScanner(os.Stdin)

	// infinite loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0{
			continue
		}

		value, ok := getCommands()[cleaned[0]]
		if ok {
			err := value.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		// fmt.Print("Your command was: ")
		// fmt.Println(cleaned[0])


	}
}

// commands here:
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, newMap := range getCommands() {
		fmt.Printf("%s: %v\n", newMap.name, newMap.description)
	}
	return nil
}



// the clean Input function
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words


}



