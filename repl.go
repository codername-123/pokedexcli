package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
		command.callback()
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
