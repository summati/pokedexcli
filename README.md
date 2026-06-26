# Pokedex CLI

A command-line Pokédex application built with Go that allows you to explore Pokémon data from the PokéAPI.

## Features

- **Interactive REPL Interface**: Explore Pokémon data through an interactive command-line interface
- **Location Browsing**: Navigate through Pokémon location areas with pagination support
- **Caching System**: Built-in caching to improve performance and reduce API calls
- **Help Command**: Get guidance on available commands

## Commands

- `help` - Display available commands
- `map` - Show the next location areas
- `mapb` - Show the previous location areas (go back)
- `exit` - Exit the application

## Building

```bash
go build
```

This will create a `pokedexcli` executable in the current directory.

## Running

```bash
./pokedexcli
```

Then use the interactive REPL to explore:

```
Pokedex > map
Pokedex > exit
```

## Project Structure

```
.
├── main.go                          # Entry point
├── repl.go                          # REPL implementation
├── command_*.go                     # Command implementations
├── internal/
│   ├── pokeapi/                     # PokéAPI client
│   │   ├── pokeapi.go               # API client setup
│   │   ├── location_area_req.go    # Location area requests
│   │   └── types_location_area.go  # Data types
│   └── pokecache/                   # Caching system
│       └── pokecache.go             # Cache implementation
├── go.mod                           # Module definition
└── Readme                           # This file
```

## Dependencies

- **Go**: 1.16 or higher
- **net/http**: Standard library HTTP client
- **encoding/json**: JSON parsing

## Implementation Details

### Caching

The application uses a time-based caching system to store API responses. Cached entries are automatically cleared after the configured interval, improving performance while keeping data relatively fresh.

### API Integration

The PokéAPI client handles all communication with the PokéAPI, including:
- Pagination support for browsing location areas
- JSON unmarshaling for structured data
- Error handling for API failures

### REPL Loop

The read-eval-print loop processes user commands and maintains application state across commands, including navigation history for browsing through location areas.

## Development

To run tests:

```bash
go test ./...
```

To check for issues:

```bash
go vet ./...
```

## License

This project is created for educational purposes.
