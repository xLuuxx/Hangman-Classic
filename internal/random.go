package game

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func RandomWord() string {
	filename := "data/words.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	words := strings.Fields(string(content)) // get the content of the file
	if len(words) == 0 {
		fmt.Println("No words found in file")
	} // Check if there is any words in the file to avoid the error for Intn.

	randomIndex := rand.Intn(len(words)) // get a random words according to the number

	return words[randomIndex]
}

func RandomLetter(word string, letters []rune) rune {
	index := rand.Intn(len(word)) // get a random letter that is in the word
	letters[index] = rune(word[index])
	return rune(word[index])
}
