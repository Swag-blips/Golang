package main

import (
	"fmt"

	"example.com/investment-calculator/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------")

	}
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing Go Bank")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
