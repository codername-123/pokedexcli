package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codername-123/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient    pokeapi.Client
	next          *string
	prev          *string
	caughtPokemon map[string]pokeapi.ResponsePokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			break
		}
		words := clearInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		var args []string
		cmdName := words[0]
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommand()[cmdName]
		if !ok {
			fmt.Println("Unknown command check help")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func clearInput(token string) []string {
	out := strings.ToLower(token)

	// splits the string on white space
	words := strings.Fields(out)
	return words
}
