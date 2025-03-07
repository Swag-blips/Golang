package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50}

	num = append(num, 30, 40, 50, 60)

	fmt.Println(num)

	slice := make([]int, 5, 20)

	fmt.Println(slice)
}
