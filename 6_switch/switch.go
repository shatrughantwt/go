package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch

	i := 2

	switch i {
	    case 1:
		fmt.Println("one")
		//break
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("other")
	}

	// Multiple switch case
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend")
	default:
		fmt.Println("Its a working day")
	}

	// Type switch case
	whoAmi := func(i interface{}) {
		switch t := i. (type){
		case int:
			fmt.Println("Its a integer")
		case bool:
			fmt.Println("Its a bool")
		case string:
			fmt.Println("Its a string")
		default:
			fmt.Println("This is ", t, "you can see")
		}
	}

	whoAmi(33.3)
}
