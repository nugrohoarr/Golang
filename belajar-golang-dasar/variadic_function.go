package main

import "fmt"

func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	numbers := []int{10, 10, 10}
	fmt.Println(Sum(numbers...))
}
