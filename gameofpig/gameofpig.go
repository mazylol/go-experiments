package gameofpig

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	turnNumber   = 1
	currentScore = 0
	roundScore   = 0
	option       string
	playAgain    string
)

func Play() {
	fmt.Println("The Game of Pig is a classic game played with a 6 sided die. In the game a player rolls the die. \nIf they roll a 2 through 6, they add that score to their round score, but if they roll a 1,\n their round is over and their round score resets to zero.")
	fmt.Println("Press enter to play...")
	_, err := fmt.Scanln()
	if err != nil {
		return
	}

start:
	for currentScore < 100 {
		playTurn()
		turnNumber++
	}

	fmt.Print("You won The Game of Pig! Play again? (yes/no) ")
	_, err = fmt.Scanln(&playAgain)
	if err != nil {
		return
	}
	if playAgain == "yes" {
		currentScore = 0
		turnNumber = 0
		goto start
	}
}

func diceRoll() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, 6
	return rand.Intn(max-min+1) + min
}

func playTurn() {
beginning:
	fmt.Printf("Turn %d\n", turnNumber)
	fmt.Printf("Your Current Score: %d\n", currentScore)
	fmt.Printf("Round Score: %d\n", roundScore)
	fmt.Print("Would you like to roll or bank? ")
	_, err := fmt.Scan(&option)
	if err != nil {
		return
	}
	fmt.Println()

	if option == "roll" {
		var num = diceRoll()
		if num == 1 {
			roundScore = 0
			fmt.Printf("You rolled a 1! You get a zero for this turn.\n\n")
			goto end
		} else {
			roundScore += num
			fmt.Printf("You rolled a %d\n\n", num)
			goto beginning
		}
	} else if option == "bank" {
		currentScore += roundScore
		roundScore = 0
	}

end:
}
