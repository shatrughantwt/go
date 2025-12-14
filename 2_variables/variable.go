package main

import "fmt"

func main() {
	// Phase 1
	// Int
	fmt.Println(35)
	// String
	fmt.Println("Boruto")
	// Boolean
	fmt.Println(true)
	// Float
	fmt.Println(34.53)

	// Phase 2
	// How to initilize a variable
	var name string = "Boruto"
	fmt.Println(name)

	// Phase 3
	var age = 20
	if age >= 18 {
		fmt.Println("i'm an adult")
	} else {
		fmt.Println("i'm a kid")
	}

	// we do not have ternary operator like js have 
}  