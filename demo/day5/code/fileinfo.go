package main

import (
	"os"
	"fmt"
)

func main()  {
	fileinfo,err := os.Stat("pw.txt")
	fmt.Println(err)
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.ModTime())
	fmt.Println(fileinfo.IsDir())
}