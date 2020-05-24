package main

import (
	"fmt"
)

func main()  {
	
	// remove 
	nums := []int{1,2,3,4,5}
	fmt.Println(nums[1:len(nums)])
	fmt.Println(nums[1:])
	fmt.Println(nums[:len(nums)-1])

	copy(nums[1:],nums[2:])  // remove 2 , middle element
	fmt.Println(nums[:len(nums)-1])
}