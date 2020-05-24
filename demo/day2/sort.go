package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{5,4,5,6,7,34,4,56}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"dfas","acv","uio","tre"}
	sort.Strings(names)
	fmt.Println(names)

	nums = []int{5,4,5,6,7,34,4,56}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	nums = []int{1,3,5,7,11,15}              // slice in order
	fmt.Println(sort.SearchInts(nums,8))     // insert 8 at position 4
	fmt.Println(nums[sort.SearchInts(nums,8)] == 8)     // true or false if 8 at position 4
	fmt.Println(nums[sort.SearchInts(nums,5)] == 5) 




}