package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	// "fmt"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// fmt.Printf("ListLocations: c = %p, c.pokeCache = %p\n", c, c.pokeCache)
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, ok := c.pokeCache.Get(url)
	if !ok {
		// url is NOT in the cache
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()
		
		dit, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		// adding into cache
		c.pokeCache.Add(url, dit)
		dat = dit
	}

	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}