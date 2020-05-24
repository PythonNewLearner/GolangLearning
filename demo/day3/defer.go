package main

import (
	"fmt"
)

//defer在函数结束的时候执行
func test() string {
	// defer 后面接函数调用
	defer func(){
		fmt.Println("defer")
	}()
	defer func(){
		fmt.Println("deferA")
	}()
	defer func(){
		fmt.Println("deferB")
	}()
	fmt.Println("main1")
	fmt.Println("main2")	
	return "test"
}

func test2(n1,n2 int)  {
	//defer函数体内不管是否发生错误都会执行
	defer func ()  {
		fmt.Println("defer")
	}()
	fmt.Println("before")
	fmt.Println(n1/n2)
	fmt.Println("after")
}
func main()  {
	fmt.Println(test())

	test2(1,0)
}