package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codername-123/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient pokeapi.Client
	next       *string
	prev       *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			break
		}
		word := clearInput(scanner.Text())
		command, ok := getCommand()[word]
		if !ok {
			fmt.Println("Unknown command check help")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func clearInput(token string) string {
	// splits the string on white space
	words := strings.Fields(token)
	if len(words) == 0 {
		return ""
	}
	return words[0]
}
