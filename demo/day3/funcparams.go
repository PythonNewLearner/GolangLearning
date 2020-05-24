package main

import (
	"fmt"
)

func calc(n1 int, n2 int, callback func(int,int) int) int {
	rt := callback(n1,n2)
	if rt >1 && rt<100{
		return rt
	}
	return -1
}

func add(a,b int) int  {
	return a+b
}

func sub(a,b int) int  {
	return a-b
}
func main()  {
	fmt.Println(calc(1,2,add))
	fmt.Println(calc(100,43,sub))	
}