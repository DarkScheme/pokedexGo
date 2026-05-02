package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	// "fmt"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	dat, ok := c.pokeCache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		dit, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}
		// adding into cache
		c.pokeCache.Add(url, dit)
		dat = dit


	}

	locationAreaResp := LocationArea{}
	err := json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	return locationAreaResp, nil
	
}


