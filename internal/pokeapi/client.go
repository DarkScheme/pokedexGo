package pokeapi

import (
	"net/http"
	"time"
	"github.com/DarkScheme/pokedexGo/internal/pokecache"
	//"fmt"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	myCache := pokecache.NewCache(cacheInterval)

	c := Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: myCache,
	}
	//fmt.Printf("NewClient: returning Client with pokeCache = %p\n", c.pokeCache)
	return c
}
