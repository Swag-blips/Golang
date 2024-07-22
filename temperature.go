package main

import "fmt"

func mainiac() {

	var temperature float64
	fmt.Print("What is the temperature value in celcius: ")
	fmt.Scan(&temperature)

	toFarenheit := (temperature * 1.8) + 32

	fmt.Println("Your farenheit value is", toFarenheit)
}
