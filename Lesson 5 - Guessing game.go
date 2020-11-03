// Guessing game
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const MAX = 100

/* Game management, include short intro before first game.
Keep track of guesses, games played, best game. Output these at the end.
Keep running until player says they do not want to play again. */

func main() {
	var games int
	var guesses int
	var bestGame int
	var doRun bool = true

	rand.Seed(time.Now().UnixNano())
	intro()

	for doRun {
		games++
		var guessesThisGame int

		guessesThisGame += gameRound()
		guesses += guessesThisGame

		if bestGame == 0 {
			l
			bestGame = guessesThisGame
		} else if guessesThisGame < bestGame {
			bestGame = guessesThisGame
		} else {
			fmt.Println("Sorry, you didn't beat your record. U SUCK.")
		}

		var awaitingResponse bool = true

		for awaitingResponse {
			fmt.Println("Do you want to play agian? ")

			var response string
			fmt.Scanln(&response)
			response = strings.ToUpper(response)

			if response == "YES" || response == "YES." {
				awaitingResponse = false
			} else if response == "NO" || response == "NO." {
				awaitingResponse = false
				doRun = false
			} else {
				fmt.Println("Please type \"Yes\" or \"No\".")
			}
		}
	}

	printResults(games, guesses, bestGame)
}

func intro() {
	fmt.Println("This is supposed to be an intro for the guessing game.")
	fmt.Println("YEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEET")
	fmt.Println()
}

func gameRound() int {
	var randNum int = getRandNum()
	var numGuesses int
	var doRun bool = true

	fmt.Println()
	fmt.Printf("I'm thinking of a number between 1 and %d...", MAX)

	for doRun {
		fmt.Print("Your guess? ")

		var playerGuess int
		fmt.Scanln(&playerGuess)
		numGuesses++

		if playerGuess == randNum {
			if numGuesses == 1 {
				fmt.Println("Congrats u got it in 1 guess")
			} else {
				fmt.Println()
				fmt.Printf("You got it right in %d guesses!", numGuesses)
				fmt.Println()
			}

			doRun = false
		} else if playerGuess > randNum {
			fmt.Println("It's lower.")
		} else if playerGuess < randNum {
			fmt.Println("It's higher")
		}
	}

	return numGuesses
}

func getRandNum() int {
	return 1 + rand.Intn(MAX-1)
}

func printResults(games int, guesses int, bestGame int) {
	fmt.Println()
	fmt.Println("Thanks for playing with a FREAKING AI.")
	fmt.Println("Overall results:")
	fmt.Printf("Total games     = %d", games)
	fmt.Println()
	fmt.Printf("Total guesses   = %d", guesses)
	fmt.Println()
	fmt.Printf("Guesses/game    = %f", float64(guesses)/float64(games))
	fmt.Println()
	fmt.Printf("Best game ever  = %d", bestGame)
}
