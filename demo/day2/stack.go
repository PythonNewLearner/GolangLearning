package main

import (
	"fmt"
)

func main()  {
	
	stack := []string{}
	stack = append(stack,"a","b","c")

	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	fmt.Println(x)
	fmt.Println(stack)

}