package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemon is not specified check help")
	}
	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you have not caught the pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats: ")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}

	return nil
}
