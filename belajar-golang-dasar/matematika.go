package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b

	fmt.Println(c)

	var i = 10
	i += 10 // i + 10

	i %= 3
	fmt.Println(i)

	// Unary Operator
	var j = 2
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
}
