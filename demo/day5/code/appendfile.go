package main

import (
	"os"
)
func main()  {
	file, err := os.OpenFile("name.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR,os.ModePerm)
	if err != nil {
		return
	}
	file.Write([]byte("abc"))

}