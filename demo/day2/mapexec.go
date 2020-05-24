package main

import (
	"fmt"
)

func main()  {
	
	var names = []string{"hong","li","xin","li","xin","li","li"}
	var votes = map[string]int{}
	for i:=0;i<len(names);i++{

		votes[names[i]] += 1
	}

	fmt.Println(votes)
}