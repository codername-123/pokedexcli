package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("command arguments are not provided correctly")
	}
	_, ok := cfg.caughtPokemon[args[0]]
	if ok {
		return fmt.Errorf("you have already caught %s", args[0])
	}
	pokemonDetails, err := cfg.pokeClient.PokemonDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetails.Name)

	//? pokemon who are easy to catch have a base experience of about 20-30
	//? so as the baseExperience increases it will become harder to catch strong pokemon because
	//? probability of getting a number lower than 40 will decrease
	var threshold int = 40

	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)

	randVal := random.Intn(pokemonDetails.BaseExperience)
	if randVal > threshold {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonDetails.Name)
	cfg.caughtPokemon[pokemonDetails.Name] = pokemonDetails
	return nil
}
