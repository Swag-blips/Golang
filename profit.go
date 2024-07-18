package main

import (
	"fmt"
)

func mainaa() {

	// var taxRate float64
	// var revenue float64
	// var expenses float64

	revenue := getUserInput("What is your revenue: ")
	taxRate := getUserInput("What is your taxRate: ")
	expenses := getUserInput("What is your expenses: ")

	EBT, EAT, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("The earnings before tax is: %.0f\nThe earnings after tax is: %.0f\nThe ratio is: %.1f", EBT, EAT, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses

	EAT := EBT - (EBT * (taxRate / 100))

	ratio := EBT / EAT

	return EBT, EAT, ratio
}
