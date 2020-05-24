package main

import (
	"os"
	"fmt"
)
func main()  {
	file,err := os.OpenFile("test.txt",os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	file.Seek(6,os.SEEK_SET)
	ctx := make([]byte,3)
	n,err := file.Read(ctx)
	fmt.Println(n,err,string(ctx))

	fmt.Println(file.Seek(0,os.SEEK_CUR))
	
}