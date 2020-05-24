package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Getwd())    //我在运行程序的时所在的路径


	fmt.Println(os.Executable()) //程序所在的路径
}
