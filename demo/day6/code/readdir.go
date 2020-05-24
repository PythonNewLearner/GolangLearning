package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	files,_ := ioutil.ReadDir(".")
	for _,file := range files {
		fmt.Println(file.Name(),file.Size())

	}
}