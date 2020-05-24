package main

import (
	"strings"
	"fmt"
)

func main()  {
	var builder	strings.Builder
	builder.Write([]byte("this is chen\n"))
	builder.WriteString("this is bai")
	fmt.Println(builder.String())
	fmt.Println(builder.Len())
	
	builder.Reset()
	fmt.Println(builder.String())
	fmt.Println(builder.Len())
}