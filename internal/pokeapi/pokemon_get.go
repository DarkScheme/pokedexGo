package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	// "fmt"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	dat, ok := c.pokeCache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		dit, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		// adding into cache
		c.pokeCache.Add(url, dit)
		dat = dit


	}

	PokemonResp := Pokemon{}
	err := json.Unmarshal(dat, &PokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return PokemonResp, nil
	
}


