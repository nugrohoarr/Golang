package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke", i)
	// }

	for con := 0; con < 10; con++ {
		if con%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", con)
	}
}
