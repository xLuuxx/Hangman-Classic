package game

import (
	"fmt"
	"strings"
)

func Play(h Hangman) {
	// For a colorful game
	var BoldYellow = "\033[1;33m"
	var Blue = "\033[34m"
	var Red = "\033[31m"
	var BoldGreen = "\033[1;32m"
	var BoldRed = "\033[1;31m"
	var Reset = "\033[0m"

	fmt.Println(BoldYellow + "Welcome on the hangman game !" + Reset)
	fmt.Println(Blue+"The word you're looking for has", len(h.Word), "letters."+Reset)
	fmt.Println()

	// Letters that will show up at the beginning
	n := len(h.Word)/2 - 1
	for i := 0; i < n; i++ {
		var randomLetter rune
		for {
			randomLetter = RandomLetter(h.Word, h.Revealed)
			if !h.Guessed[randomLetter] {
				break
			}
		}
	}

	// Play until there is no more attempts
	for h.NbAttempts > 0 {

		// The word to guess hidden behind "_"
		for _, r := range h.Revealed {
			if r == 0 {
				fmt.Print("_")
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()

		// Ask a letter to the player
		letter := AskLetter(h)

		// If the letter has already been proposed, let the player give another one
		if h.Guessed[letter] {
			fmt.Println(Red + "You have already guessed this letter" + Reset)
			continue
		}
		h.Guessed[letter] = true

		// Reveal the emplacement if it is a good letter
		if strings.ContainsRune(h.Word, letter) {
			RevealLetter(h, letter)
		} else {
			// Reduce the number of attempts if the letter isn't in word
			fmt.Printf(Red + "This letter is not in the word.\n" + Reset)
			h.NbAttempts--
			h.WrongLetters = append(h.WrongLetters, letter)
		}

		// Shows the drawing of the Hangman thanks to the number of attempts
		fmt.Println()
		DisplayHangman(h)

		// End the game when there's no more attempts or when the word is found
		if h.NbAttempts == 0 {
			fmt.Println(BoldRed+"You've ran out of attempts, the word was :", h.Word+Reset)
		} else if WordFound(h) == true {
			fmt.Println(BoldGreen+"Well done ! You've guessed the word :", h.Word+Reset)
			return
		}
		fmt.Println()
	}
	return
}

func RevealLetter(h Hangman, letter rune) {
	for i, r := range h.Word {
		if r == letter {
			h.Revealed[i] = letter
		}
	}
}

func WordFound(h Hangman) bool {
	for _, r := range h.Revealed {
		if r == 0 {
			return false
		}
	}
	return true
}
