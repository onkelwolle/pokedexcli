package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokemon in your Pokedex:")

	for _, pokemon := range cfg.pokedex {
		fmt.Println(pokemon.Name)
	}
	return nil
}