package main

import "fmt"

func main() {
	// Phase 1
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	// Phase 2
	for i := 0; i <=3; i++ {
		fmt.Println(i)
	}
}