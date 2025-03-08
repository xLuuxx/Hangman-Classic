package game

import (
	"fmt"
)

func Restart() bool {
	var restart string

	// Ask at the end of a game if the player wants to play again
	fmt.Println("Do you want to play again ? Enter yes ! ")
	_, err := fmt.Scan(&restart)
	if err != nil {
		return false
	}
	if restart == "yes" {
		return true
	}
	return false
}
