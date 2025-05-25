package pokeapi

type AreaLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Areas    []Area  `json:"results"`
}

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonList struct {
	Pokemons []PokemonEntry `json:"pokemon_encounters"`
}

type PokemonEntry struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	TypeInfo TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

func GetAreaLocations(url string) (AreaLocations, error) {
	return doGetJSON[AreaLocations](url)
}

func GetPokemonList(url string) (PokemonList, error) {
	return doGetJSON[PokemonList](url)
}

func GetPokemomn(url string) (Pokemon, error) {
	return doGetJSON[Pokemon](url)
}
