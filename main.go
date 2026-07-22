package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/fanatcx/pokedexcli/internal/pokecache"
)

type Config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
}

func main() {
	startURL := baseURL + "/location-area?limit=20"
	cfg := &Config{
		next:  &startURL,
		cache: pokecache.NewCache(5 * time.Second),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		cleaned := CleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		comm, exists := commands[cleaned[0]]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}

		if err := comm.callback(cfg); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}