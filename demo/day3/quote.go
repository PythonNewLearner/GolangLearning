package main

import (
	"fmt"
)

func main()  {
	users := make([]string,10)
	tmpUser := users
	fmt.Printf("%p,%p\n",users,tmpUser)   // same address
	fmt.Printf("%p,%p\n",&users,&tmpUser)   // different address

	//值类型 : int,float,point,array,struct

	//引用类型: slice,map,interface

	array := [3]int{}
	tmpArray := array
	tmpArray[0] = 10
	fmt.Println(array,tmpArray)  // array no change when tmpArray change


	
	b := make([]int,10)
	fmt.Println(b)
	test2(b)  //引用类型会改变b (slice是引用类型)
	fmt.Println(b)

}
func test2(s []int)  {
	s[0] = 1
}
