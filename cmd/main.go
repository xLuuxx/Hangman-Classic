package main

import "hangman_classic/internal"

func main() {
	word := game.RandomWord()
	var wrongLetters []rune

	// Initialized a game with a random word and 10 attempts
	game.Play(game.Hangman{
		Word:         word,
		Revealed:     make([]rune, len(word)),
		NbAttempts:   10,
		Guessed:      make(map[rune]bool),
		WrongLetters: wrongLetters,
	})
	if game.Restart() == true {
		main()
	}
	return
}
