package main

import (
	"io"
	"os"
)

func main()  {
	log1,_ := os.Create("test/1.log")
	log2,_ := os.Create("test/2.log")
	writer := io.MultiWriter(log1,log2,os.Stdout)
	writer.Write([]byte("hello world"))
	log1.Close()
	log2.Close()

}