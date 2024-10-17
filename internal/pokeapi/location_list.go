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

	return locationsArea, nil
}
