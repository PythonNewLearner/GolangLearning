package main

import (
	"fmt"
	"io/ioutil"
	"os"
)



func main()  {
	file,err := os.Open("multireader.go")
	if err != nil {
		return
	}
	defer file.Close()

	ctx,_ := ioutil.ReadAll(file)
	fmt.Println(string(ctx))
}