package main

import "fmt"

func commandMapF(cfg *config, args ...string) error {
	if cfg.prev != nil && cfg.next == nil {
		return fmt.Errorf("you are on the last page")
	}
	resp, err := cfg.pokeClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = resp.Next
	cfg.prev = resp.Previous

	for _, result := range resp.Results {
		fmt.Print(" - ")
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prev == nil {
		return fmt.Errorf("you are on the first page")
	}
	resp, err := cfg.pokeClient.ListLocations(cfg.prev)
	if err != nil {
		return err
	}
	cfg.next = resp.Next
	cfg.prev = resp.Previous

	for _, result := range resp.Results {
		fmt.Print(" - ")
		fmt.Println(result.Name)
	}
	return nil
}
