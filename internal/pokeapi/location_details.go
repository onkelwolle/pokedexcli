package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(locationName string) (LocationDetails, error) {
	url := baseUrl + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationDetails{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationDetails{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	locationResp := LocationDetails{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
