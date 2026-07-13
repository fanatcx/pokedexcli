package main
import (
	"os"
	"bufio"
	"fmt"
)

func main() {

	// scanner.Scan blocks then waits for ENTER
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("Pokedex > ")
		scannerBool := scanner.Scan()

		if scannerBool == true {
			userInput := scanner.Text()
			cleanedUserInput := CleanInput(userInput)
			fmt.Printf("Your command was: %s\n", cleanedUserInput[0])
		}

		if scannerBool == false {
			
			fmt.Println("END OF FILE.")
			err := scanner.Err()
			if err == nil {
				break
			} else {
				fmt.Println("Error.")
				break
			}

		}
		

	}
	
	

}

