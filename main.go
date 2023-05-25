package main

import (
	"time"

	"github.com/codername-123/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		pokeClient:    client,
		caughtPokemon: make(map[string]pokeapi.ResponsePokemon),
	}
	startRepl(cfg)
}
