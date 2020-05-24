package main

import (
	"fmt"
)

func main()  {

	func (){
		fmt.Println("call")   // call function with no name
	}()

	func (a int)  {
		fmt.Println(a)
	}(1)
	
}