package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var userOne Person
	userOne.name = "Wayne"
	userOne.age = 20
	userOne.job = "Software engineer"
	userOne.salary = 100000

	// greet("swag")

	reverseString("swag")

}

func greet(name string) {
	fmt.Println("Hello", name)
}

func reverseString(value string) string {
	splittedString := strings.Split(value, "")
	fmt.Println(len(splittedString))

	var poppedValue string

	for i := 0; i <= len(value); i++ {

		if len(splittedString) <= 0 {
			break
		}
		fmt.Println(i)
		fmt.Println("before popped", splittedString)
		poppedValue = poppedValue + splittedString[len(splittedString)-1]

		splittedString = splittedString[:len(splittedString)-1]
		fmt.Println("after popped", splittedString, poppedValue)

	}

	return poppedValue

}
