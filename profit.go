package main

import (
	"fmt"
)

func maina() {
   
	var taxRate float64
	var revenue float64
	var expenses float64

	fmt.Print("What is your revenue ")
	fmt.Scan(&revenue)

	fmt.Print("What is your expenses amount: ")
	fmt.Scan(&expenses)

	

	EBT:= revenue - expenses
	
	fmt.Println("Earnings before tax: ", EBT)

	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	EAT:= EBT - (EBT * (taxRate/100))

	fmt.Println("Earnings after tax: ", EAT)
	ratio:= EBT/EAT
	fmt.Println("Your ratio is :", ratio)
}