package main

import (
	"errors"
	"fmt"

	"github.com/jcnnll/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	if _, ok := pokedex[args[0]]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	url := pokemonEndPoint + args[0]
	pokemon, err := pokeapi.GetPokemomn(url)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.TypeInfo.Name)
	}

	return nil
}
