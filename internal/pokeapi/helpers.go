package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/jcnnll/pokedexcli/internal/pokecache"
)

var (
	cache     *pokecache.Cache
	cacheOnce sync.Once
)

func getCach() *pokecache.Cache {
	cacheOnce.Do(func() {
		cache = pokecache.NewCache(10 * time.Minute)
	})
	return cache
}

func doGetJSON[T any](url string) (T, error) {
	var result T
	c := getCach()

	// try cache
	if data, ok := c.Get(url); ok {
		if err := json.Unmarshal(data, &result); err == nil {
			return result, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	if res.StatusCode > 299 {
		return result, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	// cache response
	c.Add(url, body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
