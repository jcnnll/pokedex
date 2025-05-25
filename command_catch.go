package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jcnnll/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon target provided")
	}

	url := pokemonEndPoint + args[0]
	pokemon, err := pokeapi.GetPokemomn(url)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if throwPokeball(pokemon) {
		pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)

	return nil
}

func throwPokeball(p pokeapi.Pokemon) bool {
	initOnce.Do(func() {
		src := rand.NewSource(time.Now().UnixNano())
		random = rand.New(src)
	})

	// dificulty scale
	factor := 0.01

	probability := 1.0 / (1.0 + factor*float64(p.BaseExperience))
	hitChance := int(probability * 100)
	throw := random.Intn(100)
	caught := throw < hitChance
	return caught
}
