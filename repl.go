package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/jcnnll/pokedexcli/internal/pokeapi"
)

const (
	locationAreaEndPoint = "https://pokeapi.co/api/v2/location-area/"
	pokemonEndPoint      = "https://pokeapi.co/api/v2/pokemon/"
)

var (
	initOnce sync.Once
	random   *rand.Rand
)

type config struct {
	nextUrl     string
	previousUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var (
	commands map[string]cliCommand
	pokedex  map[string]pokeapi.Pokemon
)

func initCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location for Pokémon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokémon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect your Pokémon pokedex",
			callback:    commandPokedex,
		},
	}
}

func initPokedex() {
	pokedex = map[string]pokeapi.Pokemon{}

}

func startRepl() {
	cfg := config{}
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := commands[commandName]
		if exists {
			err := command.callback(&cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
