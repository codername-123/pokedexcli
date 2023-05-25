package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("city name not specified check help")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	pokemonReponse, err := cfg.pokeClient.ListPokemons(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonReponse.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
