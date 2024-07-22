package main

import (
	"errors"
	"fmt"
	"os"
)

func maina() {
	// validate input

	// => show error messages & exit if invalid input is provided

	// no negative numbers
	// - not 0

	// store calculated results in a file

	// var taxRate float64
	// var revenue float64
	// var expenses float64

	revenue, err1 := getUserInput("What is your revenue: ")

	taxRate, err2 := getUserInput("What is your taxRate: ")

	expenses, err3 := getUserInput("What is your expenses: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	EBT, EAT, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("The earnings before tax is: %.0f\nThe earnings after tax is: %.0f\nThe ratio is: %.1f", EBT, EAT, ratio)
	storeResults(EBT, EAT, ratio)
}

func storeResults(ebt, eat, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio:%.3f\n", ebt, eat, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number ")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses

	EAT := EBT - (EBT * (taxRate / 100))

	ratio := EBT / EAT

	return EBT, EAT, ratio
}
