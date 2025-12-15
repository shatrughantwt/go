package main

import "fmt"

func main() {
	// Nested if-else
	age := 22

	if age >=18 {
		fmt.Println("adult")
	} else if age >= 12 {
		fmt.Println("teenager")
	} else {
		fmt.Println("kid")
	}

	// && pipe-op
	var role = "admin"
	var hasPermissions = false
	
	if role == "admin" && hasPermissions {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("No Access")
	}

	// we do not have ternary op like js have so we need to use if else 
}