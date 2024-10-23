package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check if requestet url is allready in cache, if so, return the value
	if val, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	// Read respomse from body as json
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// unmarshal json into struct
	locationsArea := LocationArea{}
	err = json.Unmarshal(dat, &locationsArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationsArea, nil
}
