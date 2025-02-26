package main

import "fmt"

func main() {

	// delcaring variables with var keyword + type being inferred
	var name = "John Doe"

	// declaring variable while inferring types
	age := 1

	// manually assinging type
	var nameWithAssignedType float32 = 12.5

	// multiple variable declaration

	var one, two, three, four, five int = 1, 2, 3, 4, 5
	fmt.Println(name, age, nameWithAssignedType)
}
