package main

import "fmt"

func main() {
	var nama [3]string // array yang di tentukan panjangnya

	nama[0] = "Tri"
	nama[1] = "Nugroho"
	nama[2] = "Dwi"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{90, 80, 70}

	fmt.Println(values)

	var values1 = [...]int{ // array yang tidak di tentukan panjangnya
		90,
		80,
		70,
		80,
	}
	fmt.Println(values1)
}
