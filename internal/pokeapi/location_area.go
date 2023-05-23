package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResponseLocationArea, error) {

	url := baseUrl + "/location-area?offset=0&limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	if body, ok := c.cache.Get(url); ok {
		var locationResp ResponseLocationArea
		err := json.Unmarshal(body, &locationResp)
		if err != nil {
			return ResponseLocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	var locationResp ResponseLocationArea
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	c.cache.Add(url, body)
	return locationResp, nil
}
