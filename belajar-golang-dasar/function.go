package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World!")
}

func sayGoodBye(firstName string, lastName string) { // Function Parameter
	fmt.Println("Good Bye " + firstName + " " + lastName)
}

func getFullName(name string) string {
	hello := "Hello " + name
	return hello
}

func getName() (string, string) {
	return "Tri", "Nugroho"
}

func main() {
	sayHello()
	sayGoodBye("Tri", "Nugroho")

	result := getFullName("Tri")
	fmt.Println(result)

	fmt.Println(getFullName("makan"))
	fmt.Println(getFullName("minum"))

	firstName, lastName := getName()
	fmt.Println(firstName, lastName)

	firstName, _ = getName() // ignore return value
	fmt.Println(firstName)

}
