package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
