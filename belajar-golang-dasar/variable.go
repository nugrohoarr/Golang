package main

import "fmt"

func main() {
	var name string

	name = "Tri Nugroho"
	fmt.Println(name)

	name = "Tri"
	fmt.Println(name)

	// best practice to declare variable with type string
	nama := "Tri"
	fmt.Println(nama)

	var (
		firstName = "Tri"
		lastName  = "Nugroho"
	)
	fmt.Println(firstName + " " + lastName)

	// multiple variable
}
