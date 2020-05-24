package main

import "fmt"


const (
	packageName string = "package"
	packageMsg = "package"
)
func main()  {

	var outter string = "outter"
	fmt.Println(outter)

	{
		var inner string ="inner"
		fmt.Println(inner)
	}
}