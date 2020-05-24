package main

import (
	"io/ioutil"
	"fmt"
)
func main()  {
	ctx,err := ioutil.ReadFile("multireader.go")
	if err != nil {
		return
	}
	fmt.Println(string(ctx))
}