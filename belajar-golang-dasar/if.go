package main

import "fmt"

func main() {
	name := "Tri ari"

	if name == "Tri" {
		fmt.Println("Hello Tri")
	} else if name == "Dwi" {
		fmt.Println("Hello Dwi")
	} else {
		fmt.Println("Hi,  Boleh Kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Namaet Panjang")
	}
}
