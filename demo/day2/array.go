package main

import (
	"fmt"
)

func main()  {
	var names [3]string
	var attendance [3]bool
	var score [3]float64
	fmt.Printf("%#v\n",names)
	fmt.Printf("%#v\n",attendance)
	fmt.Printf("%#v\n",score)

	names = [...]string{"a","b","c"}
	fmt.Printf("%#v\n",names)

	testnames := [...]string{"j","q","k"}
	fmt.Printf("%#v\n",testnames)

	names1 := [3]string{2:"kk"}  //  length is 3, index of 2 is "kk"
	fmt.Printf("%#v\n",names1)

	fmt.Printf("%s\n",names[0])
	names[0] = "m"

	fmt.Println(len(names))

	for i:=0;i<len(names);i++{
		fmt.Printf("%s\n",names[i])
	}

	for k,v := range names {
		fmt.Println(k,v)
	}


	d2 := [3][2]int{1:[2]int{1,2}}    // 2-d array   3x2
	fmt.Printf("%#v\n",d2)
	fmt.Println(d2[1][0])




















}