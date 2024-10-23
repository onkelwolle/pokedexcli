package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide an location name")
	}
	location := args[0]

	locationDetails, err := cfg.pokeapiClient.LocationDetails(location)
	if err != nil {
		return err
	}

	for _, pokemons := range locationDetails.PokemonEncounters {
		fmt.Println(pokemons.Pokemon.Name)
	}
	return nil
}
