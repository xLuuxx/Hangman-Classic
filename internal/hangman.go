package game

type Hangman struct {
	Word         string
	Revealed     []rune
	NbAttempts   int
	Guessed      map[rune]bool
	WrongLetters []rune
}
