package main

import (
	"fmt"
	"bytes"
)

func main()  {
	fmt.Printf("%#v\n",[]byte("abc"))
	fmt.Println(bytes.Compare([]byte("abc"),[]byte("xyz")))
}