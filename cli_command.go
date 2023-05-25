package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Shows the next 20 locations in the pokemon world",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previos 20 locations in the pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location-area>",
			description: "It takes the name of a location area as argument and shows the list of all Pokemons in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "catches pokemon and add them to your pokedex",
			callback:    commandCatch,
		},
	}
}
