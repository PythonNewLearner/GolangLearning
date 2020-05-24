package main

import (
	"fmt"
	"strconv"
)



func main()  {
	var numstr = "3"
	v,_ := strconv.Atoi(numstr)
	fmt.Println(v)

	var numfloat = "3.32"
	vf,_ := strconv.ParseFloat(numfloat,64)
	fmt.Println(vf)

	var i = 123
	fmt.Println(strconv.Itoa(i))

	var f =12.324
	fmt.Println(strconv.FormatFloat(f,'f',5,64))


}