package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hi, this is a mini-game where I guess a number from 1 to 100. Try to guess.")
	nRec := 10
	game(nRec)
}

func game(nRec int) {
	nChan := difficulty()
	if nChan == 0 {
		print("error")
		return
	}
	fmt.Println("Number of guessing attempts:", nChan)
	nWin := generate()
	fmt.Println("Number has been generated, try to guess!")
	a := time.Now()
	win := false
	for i := 0; i < nChan; i++ {
		number := 0

		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Enter a number without symbols or letters", err)
			i--
			continue
		}
		if number == nWin {
			fmt.Println("Your guess was correct!")
			fmt.Println("Attempts spent: ", i+1)
			if nRec > i+1 {
				nRec = i + 1
			}
			fmt.Println("U record is:", nRec, "attempts")
			win = true
			break
		} else if number > nWin {
			fmt.Println("Your guess was incorrect! The hidden number is smaller. Number of guessing attempts: ")
			nChanost := nChan - (i + 1)
			fmt.Println(nChanost)
		} else {
			fmt.Println("Your guess was incorrect! The hidden number is bigger. Number of guessing attempts: ")
			nChanost := nChan - (i + 1)
			fmt.Println(nChanost)
		}
	}
	b := time.Since(a)
	fmt.Println("Time elapsed: ", b)
	if win == false {
		fmt.Println("Number:", nWin)
	}
	fmt.Println("Would you like to play again?\n    yes or no")
	playag := ""
	_, err := fmt.Scanln(&playag)
	if err != nil {
		return
	}
	playag = strings.ToLower(playag)
	if playag == "yes" || playag == "y" {
		game(nRec)
	} else {
		return
	}
}

func generate() int {
	a := rand.Intn(100)
	a++
	return a
}

func difficulty() int {
	chanceint := 0
	for {
		fmt.Println("Now enter the difficulty: Easy, Medium, Hard. (10, 5, 3 attempts)")
		chance := ""
		_, err := fmt.Scanln(&chance)
		if err != nil {
			return 0
		}
		chance = strings.ToLower(chance)

		switch chance {
		case "easy":
			chanceint = 10
		case "medium":
			chanceint = 5
		case "hard":
			chanceint = 3
		default:
			fmt.Println("Введите сложность")
			continue
		}
		break
	}
	return chanceint
}
