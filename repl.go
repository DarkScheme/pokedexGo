package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)


func startRepl() {
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

		fmt.Print("Your command was: ")
		fmt.Println(cleaned[0])
	}
}



func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words


}



