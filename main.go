package main

import (
	"time"
	"github.com/DarkScheme/pokedexGo/internal/pokeapi"
)


func main() {

	bag := map[string]pokeapi.Pokemon{}
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	cfg := &configStruct{
		pokeapiClient: pokeClient,
		caughtPokemon: bag,
	}

	startRepl(cfg)
}

