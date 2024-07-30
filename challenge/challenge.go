package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNumber := rand.IntN(100)
	fmt.Println("Your random number is: ", randomNumber)

	var guess int
	var count int

	for randomNumber != guess {
		fmt.Print("Guess a number: ")
		fmt.Scan(&guess)
		count++
		fmt.Println(count)

		if count == 5 {
			fmt.Println("maximum tries exceeded")
			break
		}

		if randomNumber > guess {
			fmt.Println("higher")
			continue
		} else if randomNumber < guess {
			fmt.Println("Lower")
			continue
		} else if randomNumber == guess {
			fmt.Println("Correct guess")
			break
		}

	}
}
