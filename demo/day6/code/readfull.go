package main

import (
	"fmt"
	"os"
	"io"
	"bytes"
)
func main()  {
	f,_ := os.Open("multireader.go")
	defer f.Close()
	buffer := bytes.NewBuffer([]byte(""))
	io.Copy(buffer,f)
	fmt.Println(buffer)
}