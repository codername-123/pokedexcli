package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(name string) (ResponsePokemon, error) {

	url := baseUrl + "/pokemon/" + name

	if body, ok := c.cache.Get(url); ok {
		var pokemonResp ResponsePokemon
		err := json.Unmarshal(body, &pokemonResp)
		if err != nil {
			return ResponsePokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponsePokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponsePokemon{}, err
	}

	var pokemonResp ResponsePokemon
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return ResponsePokemon{}, err
	}

	c.cache.Add(url, body)
	return pokemonResp, nil
}
