package main

import (
	"io"
	"os"
	"fmt"
)

func main()  {
	f1,_ :=os.Open("test/1.log")
	f2,_ := os.Open("test/2.log")
	reader := io.MultiReader(f1,f2)
	buffer := make([]byte , 5)
	for {
		n,err := reader.Read(buffer)
		if err ==io.EOF {
			break
		}
		fmt.Println(n,err,string(buffer[:n]))

	}


	f1.Close()
	f2.Close()
}