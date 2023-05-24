package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("city name not specified check help")
	}
	pokemonReponse, err := cfg.pokeClient.ListPokemons(args[0])
	if err != nil {
		return err
	}
	for _, pokemon := range pokemonReponse.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
