package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Wellcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Nugroho Gagah", blackList)

	registerUser("admin", func(name string) bool {
		return name == "anjing"
	})
}
