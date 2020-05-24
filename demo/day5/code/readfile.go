package main

import (
	"fmt"
	"os"
	"io"
)

func main()  {


	file, err := os.Open("test.txt")
	if err != nil{
		return
	}
	defer file.Close()
	cxt:= make([]byte,10)
	for {
		_,err := file.Read(cxt)
		if err == io.EOF{
			break
		}
		
		fmt.Println(string(cxt))
	}
	


}