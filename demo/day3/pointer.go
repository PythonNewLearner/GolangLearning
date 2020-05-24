package main

import (
	"fmt"
)


func change(value int)  {
	value += 1
}
//值类型可通过指针改变其值
func changePtr(ptr *int)  {
	*ptr = *ptr+2
}

func main()  {
	value := 1
	change(value)
	fmt.Println(value)

	changePtr(&value)  //传指针，改变值类型变量
	fmt.Println(value)
	
}