package main

import (
	"fmt"
)

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	//

	for counter := 0; counter < 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Selesai")

	names := []string{"Tri", "Nugroho", "Yosef", "Irawan", "Dwi", "Wahyudi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range names {
		fmt.Println(value)
	}

	fmt.Println(factorialLoop(10))
	// for i, value := range names {
	// 	fmt.Println("Index", i, "=", value)
	// }
	// for _, value := range names {
	// 	fmt.Println("Value", value)
	// }
}
