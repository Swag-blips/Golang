package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	principal, err := getUserInputAlt("What is your Principal: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	interestRate, err1 := getUserInputAlt("What is your Interest rate (in %): ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	numberOfTimes, err2 := getUserInputAlt("What is the number of times interest is compounded per year: ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	numberOfYears, err3 := getUserInputAlt("What is the number of years: ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	compoundInterest := calculateCompoundInterest(principal, interestRate, numberOfTimes, numberOfYears)
	fmt.Printf("The future value of the investment is: %.2f\n", compoundInterest)
}

func getUserInputAlt(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input must be greater than 0")
	}

	return userInput, nil
}

func calculateCompoundInterest(principal, interestRate, numberOfTimes, numberOfYears float64) float64 {
	ratePerTime := (interestRate / 100) / numberOfTimes
	exponent := numberOfTimes * numberOfYears
	result := principal * math.Pow((1+ratePerTime), exponent)

	return result
}
