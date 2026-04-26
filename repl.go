package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type cliCommand struct {
	name string
	description string
	callback func(*configStruct) error
}

type configStruct struct {
	Next string
	Previous string
}

type locationAreasResp struct {
	Count int
	Next string
	Previous string
	Results []struct{
		Name string
		URL string
	}
}

func getCommands() map[string]cliCommand {
	// creates the command map
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
					"map": {
						name: "map",
						description: "Displays the next 20 locations",
						callback: commandMap,
					},
					"mapb": {
						name: "mapb",
						description: "Displays the previous 20 locations",
						callback: commandMapb,
					},
				}
	return newMap
}


func startRepl() {

	// creates a buffer scanner
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &configStruct{}

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
			err := value.callback(cfg)
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
func commandExit(c *configStruct) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *configStruct) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, newMap := range getCommands() {
		fmt.Printf("%s: %v\n", newMap.name, newMap.description)
	}
	return nil
}

func commandMap(c *configStruct) error {
	// this is the get request
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	} 
	res, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil{
		log.Fatal(err)
	}
	// fmt.Printf("%s", body)

	// and this is the Unmarshal JSON to struct

	area := locationAreasResp{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}

	c.Next = area.Next
	c.Previous = area.Previous

	for _, a := range area.Results {
		fmt.Println(a.Name)
	}
	
	return nil
}

func commandMapb(c *configStruct) error {
	// this is the get request
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Previous != "" {
		url = c.Previous
	} 
	res, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil{
		log.Fatal(err)
	}
	// fmt.Printf("%s", body)

	// and this is the Unmarshal JSON to struct

	area := locationAreasResp{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}

	c.Next = area.Next
	c.Previous = area.Previous

	for _, a := range area.Results {
		fmt.Println(a.Name)
	}
	
	return nil
}



// the clean Input function
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words


}



