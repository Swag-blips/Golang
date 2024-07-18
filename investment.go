package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func maina() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputTexta("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputTexta("How many years?")
	fmt.Scan(&years)

	outputTexta("What is your expected return rate? ")
	fmt.Scan(&expectedReturnRate)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("The value (inflaftion adjusted is): %.0f\n", futureRealValue)
	// fmt.Println("The future value is :", futureValue)
	fmt.Printf(`Future Value: %.1f
	The value (inflaftion adjusted is) : %.0f`, futureValue, futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)
}

func outputTexta(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv

	return fv, rfv
}
