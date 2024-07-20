package main

import "fmt"

func mainaaaa() {

	for i := 0; i < 100; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
