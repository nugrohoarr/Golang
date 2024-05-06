package main

import "fmt"

func main() {
	name := "Tri"

	switch name {
	case "Tri":
		fmt.Println("Hello Tri")
	case "Dwi":
		fmt.Println("Hello Dwi")
	default:
		fmt.Println("Hi,  Boleh Kenalan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah panjang")
	}
}
