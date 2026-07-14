package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// first run
	for {
		fmt.Print("Pokedex > ")

		// scanner returned false(possible error)
		if !scanner.Scan() {
			break
		}
		cleaned := CleanInput(scanner.Text())

		// check for empty ENTER from the user, move to next command
		if len(cleaned) == 0 {
			continue
		} 
		// check for command first
		if len(cleaned) == 1 {
			com, exist := commands[cleaned[0]]
			if exist {
				if err := com.callback(); err != nil {
					fmt.Println("Error: ", err)
				}
			} else {
				// could be more verbose
				fmt.Println("Uknown command or invalid pokemon")
				continue

			}
			// does not exist continue or exit. for future code
			continue 
		}

		// the default reply which will search for the pokemon. for future code
		fmt.Println("Future Reply")
		continue

		}

		if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	}

	

