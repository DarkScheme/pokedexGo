package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	// "io"
	// "log"
	// "net/http"
	// "encoding/json"
	"github.com/DarkScheme/pokedexGo/internal/pokeapi"
	"errors"
	"math/rand"

)

type cliCommand struct {
	name string
	description string
	callback func(*configStruct, []string) error
}

type configStruct struct {
	pokeapiClient pokeapi.Client
	Next *string
	Previous *string
	caughtPokemon map[string]pokeapi.Pokemon
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
					"explore": {
						name: "explore",
						description: "explores a given are. E.g. explore canalave-city-area",
						callback: commandExplore,
					},
					"catch": {
						name: "catch",
						description: "catches the pokemon (duh.). Need to write the pokemon name, e.g.: catch pikachu",
						callback: commandCatch,
					},
					"inspect": {
						name: "inspect",
						description: "inspects the pokemon. Only works for already caught Pokemon, e.g.: inspect pikachu",
						callback: commandInspect,
					},
					"pokedex": {
						name: "pokedex",
						description: "shows your Pokedex (all Pokemon that you have caught so far)",
						callback: commandPokedex,
					},
				}
	return newMap
}


func startRepl(c *configStruct) {

	// creates a buffer scanner
	scanner := bufio.NewScanner(os.Stdin)

	// infinite loop
	for {
		fmt.Println(" ")
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0{
			continue
		}

		value, ok := getCommands()[cleaned[0]]
		if ok {
			err := value.callback(c, cleaned[1:])
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

// the clean Input function
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words


}

// commands here:
func commandExit(c *configStruct, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *configStruct, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, newMap := range getCommands() {
		fmt.Printf("%s: %v\n", newMap.name, newMap.description)
	}
	return nil
}

func commandMap(cfg *configStruct, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
	// // *the below is old code, keeping for now just in case*
	// // // this is the get request
	// url := "https://pokeapi.co/api/v2/location-area/"
	// if c.Next != "" {
	// 	url = c.Next
	// } 
	// res, err := http.Get(url)
	
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// // // fmt.Printf("%s", body)

	// // // and this is the Unmarshal JSON to struct

	// area := locationAreasResp{}
	// err = json.Unmarshal(body, &area)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// c.Next = area.Next
	// c.Previous = area.Previous

	// for _, a := range area.Results {
	// 	fmt.Println(a.Name)
	// }
	
	// return nil
}

func commandMapb(cfg *configStruct, args []string) error {

	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
	// // *the below is old code, keeping for now just in case*
	// // this is the get request 
	// url := "https://pokeapi.co/api/v2/location-area/"
	// if c.Previous != "" {
	// 	url = c.Previous
	// } 
	// res, err := http.Get(url)
	
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// // fmt.Printf("%s", body)

	// // and this is the Unmarshal JSON to struct

	// area := locationAreasResp{}
	// err = json.Unmarshal(body, &area)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// c.Next = area.Next
	// c.Previous = area.Previous

	// for _, a := range area.Results {
	// 	fmt.Println(a.Name)
	// }
	
	// return nil







// explore command
func commandExplore(cfg *configStruct, args []string) error {

	if len(args) <= 0 {
		return errors.New("for this command you must also specify the area name")
	}
	areaName := args[0]

	locationsArea, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	message := "Exploring " + areaName + "..."
	fmt.Println(message)
	fmt.Println("Found Pokemon:")

	for _, pok := range locationsArea.PokemonEncounters {
		pokis := "- " + pok.Pokemon.Name
		fmt.Println(pokis)
	}
	return nil

	
}


// catch command
func commandCatch(cfg *configStruct, args []string) error {

	if len(args) <= 0 {
		return errors.New("for this command you must also specify the pokemon name")
	}

	catchphrase := "Throwing a Pokeball at " + args[0] + "..."
	fmt.Println(catchphrase)

	poke, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	getBelow := 75
	roll := rand.Intn(poke.BaseExperience)
	if roll < getBelow {
		cfg.caughtPokemon[poke.Name] = poke
		fmt.Printf("%s was caught!\n", poke.Name)
	} else {
		fmt.Printf("%s escaped!\n", poke.Name)
	}


	return nil
}

// inspect command
func commandInspect(cfg *configStruct, args []string) error {

	if len(args) <= 0 {
		return errors.New("for this command you must also specify the pokemon name")
	}

	pokName := args[0]

	poki, ok := cfg.caughtPokemon[pokName]
	if !ok {
		fmt.Println("Inspection not possible, you have not caught this Pokemon")
		return nil
	}
	
	fmt.Printf("Height: %d\n", poki.Height)
	fmt.Printf("Weight: %d\n", poki.Weight)
	fmt.Printf("BaseExperience: %d\n", poki.BaseExperience)
	fmt.Println("Stats:")
	for _, s := range poki.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poki.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	


	return nil


}


func commandPokedex(cfg *configStruct, args []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", key)
	}
	return nil
}

