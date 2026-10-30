package prac

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random no between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I Have choosed a no between 1 and 100")
	fmt.Println("Can you guess the no?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			println("Congratulation! you guessed the correct no!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher no.")
		} else {
			fmt.Println("Guess is too high try to guess lower no")
		}
	}
}
