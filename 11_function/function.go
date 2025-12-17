package main

import (
	"fmt"
)

func add(a,b int)int {
	return a+b
}

func getLang()(string,string,string,bool){
	return "boruto","go","js",true
}

func applyOp(n int, operation func(int)int)int{
	return operation(n)
}

func multiplier(factorial int) func(int)int{
	return func(n int) int{
		return n * factorial
	}
}
func main()  {
	result := add(2,3)
	fmt.Println(result)

	lang1, lang2, lang3, isOk := getLang()

	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)
	fmt.Println(isOk)
	double := func (x int) int {
		return x*2
	}

	square := func(x int) int {
		return x*x
	}

	fmt.Println(applyOp(5, double))
	fmt.Println(applyOp(5, square))

	timesTwo := multiplier(2)
	timesThree := multiplier(3)

	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
}