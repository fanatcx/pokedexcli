package main

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	previous *string 
	next *string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var config Config 

	for {
		fmt.Print("Pokedex > ")
		// handled after loop
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
		// command fails
		if err := comm.callback(&config); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	
	// specific scanner err
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

}


	

