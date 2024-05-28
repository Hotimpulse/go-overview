package main

import "fmt"

func main() {
	var name string
	name = "Lev"
	number := 9

	// var bl bool
	// fmt.Println(bl)
	fmt.Println("Hello World!", name, number)

	fmt.Printf(("Enter your name: "))

	// fmt.Printf("%T", number)

	fmt.Scan(&name) // ampersand allocates a place in memory for the variable

	fmt.Println(name)

	fmt.Printf("Hello, %v, welcome to the app! \n", name)
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	// fmt.Println(age >= 10) // returns a boolean

	if age >= 10 {
		fmt.Println("Good job!")
	} else {
		fmt.Println("Sorry, you can't use this app!")
	}
}
