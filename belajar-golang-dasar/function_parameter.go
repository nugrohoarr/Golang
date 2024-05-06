package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filterName := filter(name)
	fmt.Println("Hello", filterName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Nugroho", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
