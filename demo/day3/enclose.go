package main

import (
	"fmt"
)

func addBase(base int) func(int) int  {    // func addBase return a func
	return func(n int) int{
		return n +base
	}
}

func main()  {
	add1 := addBase(1)
	fmt.Printf("%#v\n",add1)
	fmt.Println(add1(5))
}


