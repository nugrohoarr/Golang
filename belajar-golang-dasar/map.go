package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Tri",
		"city": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])

	books := make(map[string]string)
	books["title"] = "Belajar Golang"
	books["author"] = "Tri Nugroho"
	books["Ups"] = "Salah"

	fmt.Println(books)

	delete(books, "Ups")

	fmt.Println(books)
}
