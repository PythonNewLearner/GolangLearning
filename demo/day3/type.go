package main

import (
	"fmt"
)


func test(a int,b string)  {
	fmt.Println(a,b)
}

func add(a,b int) int  {
	return a+b
}
func main()  {
	fmt.Printf("%T\n",test)
	a := test
	fmt.Printf("%T\n",a)
	a(1,"a")

	var callback func(int,int) int     //函数类型变量
	fmt.Printf("%T, %#v\n",callback,callback)

	callback = add    // 同一类型才能赋值
	fmt.Println(callback(1,3))  


}