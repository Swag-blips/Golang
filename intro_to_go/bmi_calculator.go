package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	weight, err := userInput("What is your weight (in kg): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	height, err := userInput("What is your height (in meters): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	bmi := calculateBMI(weight, height)
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	printToFile(weight, height, bmi)
}

func userInput(info string) (float64, error) {
	var userInput float64
	fmt.Print(info)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("please input a positive number")
	}
	return userInput, nil
}

func calculateBMI(weight, height float64) float64 {
	return weight / math.Pow(height, 2)
}

func printToFile(weight, height, bmi float64) {
	results := fmt.Sprintf("Weight: %.2f kg\nHeight: %.2f m\nBMI: %.2f\n", weight, height, bmi)
	err := os.WriteFile("bmi.txt", []byte(results), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
