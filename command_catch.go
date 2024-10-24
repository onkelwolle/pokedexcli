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

	chanceToCatch := rand.Intn(300)

	fmt.Printf("Chance to catch: %d\n", chanceToCatch)
	fmt.Printf("Base experience: %d\n", pokemon.BaseExperience)

	if chanceToCatch >= pokemon.BaseExperience {
		fmt.Println("You caught it!")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println("You didn't catch it :(")
	}

	return nil
}
