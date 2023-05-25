package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range getCommand() {
		fmt.Printf("%s:\n	- %s\n\n", v.name, v.description)
	}

	fmt.Println()
	return nil
}
