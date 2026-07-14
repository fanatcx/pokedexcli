package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

		fmt.Printf("Your command was: %s\n", cleaned[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}