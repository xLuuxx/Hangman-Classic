package game

import (
	"fmt"
)

func AskLetter(h Hangman) rune {
	var guess string
	var Magenta = "\033[35m"
	var Blue = "\033[34m"
	var Reset = "\033[0m"
	var Red = "\033[31m"

	for {
		// Information on the state of the game before asking for a letter
		fmt.Printf("Attempts remaining : %d\n", h.NbAttempts)
		fmt.Printf(Magenta+"Letters that are not in the word : %s\n", string(h.WrongLetters)+Reset)
		fmt.Println()
		fmt.Print(Blue + "Enter a letter: " + Reset)
		_, err := fmt.Scan(&guess)
		if err != nil {
			return 0
		}
		// Check if it is only one letter
		if len(guess) == 1 && guess[0] >= 'a' && guess[0] <= 'z' {
			return rune(guess[0])
		}
		fmt.Println(Red + "Please enter a valid letter." + Reset)
	}
}
