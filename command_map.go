package main

import (
	"fmt"

	"github.com/jcnnll/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config, args ...string) error {
	url := locationAreaEndPoint

	if cfg.nextUrl != "" {
		url = cfg.nextUrl
	}

	locations, err := pokeapi.GetAreaLocations(url)
	if err != nil {
		return err
	}

	for _, area := range locations.Areas {
		fmt.Println(area.Name)
	}

	if locations.Next != nil {
		cfg.nextUrl = *locations.Next
	} else {
		cfg.nextUrl = ""
	}
	if locations.Previous != nil {
		cfg.previousUrl = *locations.Previous
	} else {
		cfg.previousUrl = ""
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := locationAreaEndPoint

	if cfg.previousUrl != "" {
		url = cfg.previousUrl
	}

	locations, err := pokeapi.GetAreaLocations(url)
	if err != nil {
		return err
	}

	for _, area := range locations.Areas {
		fmt.Println(area.Name)
	}

	if locations.Next != nil {
		cfg.nextUrl = *locations.Next
	} else {
		cfg.nextUrl = ""
	}
	if locations.Previous != nil {
		cfg.previousUrl = *locations.Previous
	} else {
		cfg.previousUrl = ""
	}
	return nil
}
