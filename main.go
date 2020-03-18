package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number!")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100) // generates numbers between 0 and n (100)

	var guess int
	for {
		fmt.Println("Please input your guess: ")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too Big")
		} else if guess < secretNumber {
			fmt.Println("Too Small")
		} else {
			fmt.Println("You win!!! ")
			break
		}
	}

}
