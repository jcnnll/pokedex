package main

import (
	"errors"
	"fmt"

	"github.com/jcnnll/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no area location provided")
	}
	url := locationAreaEndPoint + args[0]

	pokemonList, err := pokeapi.GetPokemonList(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", args[0])
	for _, pokemon := range pokemonList.Pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
