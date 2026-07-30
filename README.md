# Pokédex CLI

A Pokédex that lives in your terminal. Explore the Pokémon world, catch Pokémon, and inspect their stats — all through a REPL powered by [PokéAPI](https://pokeapi.co/), written in Go.

```
Pokedex > catch dragonite
Throwing a Pokeball at dragonite...
Congratulations! dragonite has been caught and added to the pokedex!
You may now inspect it with the inspect command!
```

Bonus: a sleeping Snorlax sees you off when you exit. 💤

## Quick Start

**Requirements:** Go 1.22+

```bash
git clone https://github.com/fanatcx/pokedexcli.git
cd pokedexcli
go run .
```

Or build a binary:

```bash
go build -o pokedexcli
./pokedexcli
```

You'll land in the REPL:

```
Pokedex > help
```

## Commands

| Command | Usage | Description |
|---|---|---|
| `help` | `help` | Display all available commands |
| `map` | `map` | List the next 20 location areas |
| `mapb` | `mapb` | List the previous 20 location areas |
| `explore` | `explore <area-name>` | List all Pokémon found in a location area |
| `catch` | `catch <pokemon-name>` | Throw a Pokéball — you have a chance to catch it! |
| `inspect` | `inspect <pokemon-name>` | View the stats, types, height, and weight of a caught Pokémon |
| `pokedex` | `pokedex` | Display every Pokémon you've caught |
| `exit` | `exit` | Close the Pokédex (Snorlax says goodbye) |

### Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area area...
tentacool
magikarp
gyarados
...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
Congratulations! magikarp has been caught and added to the pokedex!

Pokedex > inspect magikarp

Name: magikarp
Height: 9
Weight: 100
Stats:
	-hp: 20
	-attack: 10
	...
Types:
	- water
```

## How It Works

- **REPL loop** — reads input from stdin, normalizes it (lowercase, whitespace-trimmed), and dispatches to a command registry via callback functions.
- **PokéAPI integration** — location areas, encounters, and Pokémon data are fetched live over HTTP and unmarshaled into typed Go structs.
- **In-memory cache** (`internal/pokecache`) — API responses are cached by URL with a mutex-guarded map. A background goroutine (`reapLoop`) runs on a ticker and evicts entries older than the configured interval, so repeated `map` / `mapb` / `explore` calls don't hammer the API.
- **Your Pokédex** — caught Pokémon are stored in memory for the session and browsable with `pokedex` and `inspect`.

## Project Structure

```
pokedexcli/
├── main.go                      # Entry point, REPL loop, Config
├── commands.go                  # Command registry + all command callbacks
├── repl.go                      # Input cleaning
├── repl_test.go                 # Tests for input handling
├── types.go                     # PokéAPI response structs
└── internal/
    └── pokecache/
        └── pokecache.go         # Thread-safe TTL cache with background reaper
```

## Future Endeavors

- [ ] **ASCII sprites in the terminal** — render each Pokémon's sprite (PokéAPI serves sprite PNGs) as ASCII/ANSI art inside `inspect` and `catch`. May be hard. Will be worth it.
- [ ] **Dynamic catch rates** — scale catch difficulty off each Pokémon's base experience instead of a flat 1-in-3, so a Caterpie is a gimme and a Dragonite makes you sweat.
- [ ] **Persistent Pokédex** — save caught Pokémon to a JSON file on exit and load them on startup, so your collection survives between sessions.

## License

See [LICENSE](LICENSE).
