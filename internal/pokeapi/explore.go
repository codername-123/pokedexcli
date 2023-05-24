package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(city string) (ResponseExplore, error) {

	url := baseUrl + "/location-area/" + city

	if body, ok := c.cache.Get(url); ok {
		var exploreResp ResponseExplore
		err := json.Unmarshal(body, &exploreResp)
		if err != nil {
			return ResponseExplore{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponseExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseExplore{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseExplore{}, err
	}

	var exploreResp ResponseExplore
	err = json.Unmarshal(body, &exploreResp)
	if err != nil {
		return ResponseExplore{}, err
	}

	c.cache.Add(url, body)
	return exploreResp, nil
}
