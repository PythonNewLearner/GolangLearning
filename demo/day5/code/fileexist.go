package main

import (
	"os"
	"fmt"
)

func FileIsExists(path string) bool  {
	_, err := os.Stat(path)
	if err == nil {
		
		return true
	}else if os.IsNotExist(err){
		return false
	}else {
		panic(err)
	}
	
}

func main()  {
	fmt.Println(FileIsExists("pw.txt"))
	
}