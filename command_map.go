package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	locationsAreas, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsAreas.Next
	cfg.prevLocationsURL = locationsAreas.Previous

	for _, location := range locationsAreas.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("no previous locations")
	}

	locationsAreas, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsAreas.Next
	cfg.prevLocationsURL = locationsAreas.Previous

	for _, location := range locationsAreas.Results {
		fmt.Println(location.Name)
	}
	return nil
}
