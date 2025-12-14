package main

import "fmt"

const age = 40

// const group
const(
	PORT = 3000
	URL = "http://localhost"
)

func main() {

	// short hand syntax of const
    name := "boruto" // we cannot use short hand syntax outside of the func main
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(PORT)
	fmt.Println(URL)

}