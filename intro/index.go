package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {

	address := 10

	var pointerToX *int = &address

	fmt.Println(*pointerToX)
}
