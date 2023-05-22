package main

import (
	"time"

	"github.com/codername-123/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient: client,
	}
	startRepl(cfg)
}
