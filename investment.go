package main

import (
	"fmt"
	"math"
)

func main	() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("How many years? ")
	fmt.Scan(&years)

	fmt.Print("What is your expected return rate? ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("The future value is :", futureValue)
	fmt.Printf("Future Value: %.1f\nThe value (inflaftion adjusted is) : %.0f", futureValue, futureRealValue)

}
