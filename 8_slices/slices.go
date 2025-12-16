package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums = make([]int,0,5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// fmt.Println("------")


	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	var nums2 = make([]int,len(nums))
	copy(nums2,nums)
	// fmt.Println(nums2)
	// fmt.Println(nums)

	var numx = []int{1,2,3}
	var numy = []int{1,2,3}

	// Slice package
	fmt.Println(slices.Equal(numx,numy))

	var nums2D = [] []int{{1,2}, {3,4}}
	fmt.Println(nums2D)

	// Slice operator

	var nums3 = []int{0,1,2,3}
	fmt.Println(nums3[1:4])
}