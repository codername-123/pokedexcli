package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range getCommand() {
		fmt.Printf("%s: %s\n", k, v.description)
	}

	fmt.Println()
	return nil
}
