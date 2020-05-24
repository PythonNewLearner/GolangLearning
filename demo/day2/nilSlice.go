package main

import (
	"fmt"
)

func main()  {
	var nilSlice []int  // nil slice
	fmt.Printf("%#v\n",nilSlice)

	var emptySlice []int = []int{}  // emptySlice := []int{}
	fmt.Printf("%#v\n",emptySlice)

	nilSlice = append(nilSlice,1)
	emptySlice = append(emptySlice,1)
	fmt.Printf("%#v, %#v\n",nilSlice,emptySlice)
}