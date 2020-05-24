package main

import (
	"fmt"
)

func main()  {
	var ptr *int    // nil

	fmt.Printf("%T %v\n",ptr,ptr)

	i := 11
	ptr = &i
	fmt.Printf("%v\n",ptr)

	var ptr1 = new(string)   //  make a pointer using new
	fmt.Printf("%T %#v\n",ptr1,ptr1)

	var ptr2 = &ptr
	fmt.Printf("%T %#v\n",ptr2,ptr2)
}