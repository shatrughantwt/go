package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "boruto"
	m["x"] = "twt"
	m["github"] = "borutotwt"

	// fmt.Println(m["name"])
	// fmt.Println(m["phone"])
	
	// map2 := make(map[string]int)
	
	// map2["age"] = 40
	
	// fmt.Println(len(map2))
	
	
	
	// fmt.Println(m)
	// delete(m, "x")
	// fmt.Println(m)
	// clear(m)
	// fmt.Println(m)

	map3 := map[string]int{"price": 50, "phone": 234}
	fmt.Println(map3)

	v, ok := map3["x"]
	fmt.Println(v)
	if ok{
		fmt.Println("price exist")
	} else {
		fmt.Println("Not sure")
	}

	map4 := map[string]int{"price": 50, "x": 34}
	map5 := map[string]int{"price": 50, "x":34}

	fmt.Println(maps.Equal(map4, map5))
}