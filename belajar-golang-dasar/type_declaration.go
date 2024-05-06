package main

import "fmt"

func main() {
	type NoKtp string

	var ktp NoKtp = "123434234"

	var contoh string = "1231232"

	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktp)
	fmt.Println(contohKtp)

}
