package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("pokemon name missing")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon: ", pokemonName)

	// catch the pokemon
	const threshold = 50
	randNum := rand.IntN(pokemon.BaseExperience)

	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("%s ran away, try again later!\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("Congratulations you have caught %s\n", pokemonName)

	return nil
}
