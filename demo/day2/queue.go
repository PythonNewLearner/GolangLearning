package main

import (
	"fmt"
)
func main()  {
	queue := []string{}

	queue = append(queue,"a","b")
	queue = append(queue,"c")
	x := queue[0]
	queue = queue[1:]
	fmt.Println(x)
	fmt.Println(queue)


	x = queue[0]
	queue = queue[1:]
	fmt.Println(x)
	fmt.Println(queue)
	
	




}