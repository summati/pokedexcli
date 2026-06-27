package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("location name missing")
	}

	locationName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Location Area: ", locationName)

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" -%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
