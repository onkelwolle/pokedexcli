package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.PokemonDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Catching %s!\n", pokemon.Name)

	chanceToCatch := rand.Intn(pokemon.BaseExperience)

	if chanceToCatch <= 40 {
		fmt.Println("You caught it!")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println("You didn't catch it :(")
	}

	return nil
}
