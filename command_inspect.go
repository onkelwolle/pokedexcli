package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("must provide a pokemon name")
	}

	if pokemon, ok := cfg.pokedex[args[0]]; ok {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Println("height: " + fmt.Sprint(pokemon.Height))
		fmt.Println("weight: " + fmt.Sprint(pokemon.Weight))
		fmt.Println("Stats:")

		for _, stat := range pokemon.Stats {
			fmt.Println("  -" + stat.Stat.Name + ": " + fmt.Sprint(stat.BaseStat))
		}

		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Println("  -" + typ.Type.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon yet")
	}
	return nil
}