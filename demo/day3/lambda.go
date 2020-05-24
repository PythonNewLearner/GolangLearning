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

func main()  {


	add := func (a,b int) int  {
		return a+b
	}

	sub := func (a,b int) int  {
		return a-b
	}

	fmt.Println(calc(1,2,add))
	fmt.Println(calc(100,43,sub))	

	rt := calc(50,30,func(n1,n2 int) int {    // lambda function: "func(n1,n2 int) int { return n1 + n2}" is param
		return n1 + n2
	})
	fmt.Println(rt)
}