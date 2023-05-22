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
		cmd := clearInput(scanner.Text())
		fmt.Println(cmd)
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
