# Pokedex CLI

A command-line Pokédex application built with Go that lets you browse Pokémon location areas, explore encounters, and catch Pokémon from the PokéAPI.

## Features

- Interactive REPL experience for entering commands
- Browse Pokémon location areas with pagination using `map` and `mapb`
- Explore a specific location area with `explore`
- Catch Pokémon and store them in your local Pokédex
- Inspect caught Pokémon details such as stats, abilities, and types
- Built-in HTTP caching to reduce repeated API requests

## Requirements

- Go 1.26.4 or newer
- Internet access for fetching data from the PokéAPI

## Quick start

Clone the repository and build the binary:

```bash
git clone https://github.com/summati/pokedexcli.git
cd pokedexcli
go build -o pokedexcli .
```

Run it locally:

```bash
./pokedexcli
```

To install it system-wide on macOS or Linux:

```bash
sudo mv pokedexcli /usr/local/bin/
pokedexcli
```

## Available commands

- `help` - Show the available commands
- `map` - Show the next page of location areas
- `mapb` - Show the previous page of location areas
- `explore {location_name}` - List Pokémon found in a location area
- `catch {pokemon_name}` - Try to catch a Pokémon
- `inspect {pokemon_name}` - Show details for a caught Pokémon
- `pokedex` - List the Pokémon currently stored in your Pokédex
- `exit` - Quit the app

## Example session

```text
$ ./pokedexcli
> help
> map
> explore kanto
> catch pikachu
> inspect pikachu
> pokedex
> exit
```

## Project structure

```text
.
├── main.go                          # Application entry point
├── repl.go                          # REPL command loop
├── command_*.go                     # Command handlers
├── internal/
│   ├── pokeapi/                     # PokéAPI client and response models
│   │   ├── pokeapi.go
│   │   ├── location_area_req.go
│   │   ├── pokemon_req.go
│   │   └── types_location_area.go
│   └── pokecache/                   # Time-based HTTP response cache
│       └── pokecache.go
├── go.mod                           # Module definition
└── README.md                        # Project documentation
```

## Development

Run the test suite:

```bash
go test ./...
```

Run the linter/static checks:

```bash
go vet ./...
```

## Notes

- The CLI depends on the public PokéAPI, so it requires a working network connection.
- Pokémon capture is randomized, so the same attempt may not succeed every time.

## License

This project is intended for educational purposes.
