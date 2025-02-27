package main

import (
	"fmt"
	"strings"
)

const user = "admin"

func main() {

	var num1 = 2000

	fmt.Println(num1 + 2)
	// delcaring variables with var keyword + type being inferred
	var name = "John Doe"

	// declaring variable while inferring types
	age := 1

	// manually assinging type
	var nameWithAssignedType float32 = 12.5

	// multiple variable declaration

	fmt.Println(name, age, nameWithAssignedType)

	fmt.Println(strings.Compare(name, user))

}
